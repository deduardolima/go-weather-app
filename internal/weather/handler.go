package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationService interface {
	GetLocation(cep string) (string, error)
}

type WeatherService interface {
	GetWeather(city string) (float64, error)
}

type WeatherHandler struct {
	locationService LocationService
	weatherService  WeatherService
}

func NewWeatherHandler(locationService LocationService, weatherService WeatherService) *WeatherHandler {
	return &WeatherHandler{
		locationService: locationService,
		weatherService:  weatherService,
	}
}

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := h.locationService.GetLocation(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	fmt.Println("City:", city)

	tempC, err := h.weatherService.GetWeather(city)
	if err != nil {
		http.Error(w, "failed to get weather data", http.StatusInternalServerError)
		return
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273.15

	response := WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
