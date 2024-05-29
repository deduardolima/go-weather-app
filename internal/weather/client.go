package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type WeatherClient struct{}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{}
}

func (wc *WeatherClient) GetWeather(location string) (float64, error) {
	apiKey := viper.GetString("WEATHER_API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("WEATHER_API_KEY not set")
	}
	encodedLocation := url.QueryEscape(location)
	apiURL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, encodedLocation)
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to get weather data")
	}

	var result WeatherAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}
	return result.Current.TempC, nil
}
