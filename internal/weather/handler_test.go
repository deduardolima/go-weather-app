package weather

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockLocationService struct{}

func (mls *MockLocationService) GetLocation(cep string) (string, error) {
	if cep == "80010100" {
		return "Curitiba", nil
	}
	return "", fmt.Errorf("can not find zipcode")
}

type MockWeatherService struct{}

func (mws *MockWeatherService) GetWeather(city string) (float64, error) {
	if city == "Curitiba" {
		return 19.0, nil
	}
	return 0, fmt.Errorf("failed to get weather data")
}

func TestWeatherHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/weather?cep=80010100", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	locationService := &MockLocationService{}
	weatherService := &MockWeatherService{}
	handler := NewWeatherHandler(locationService, weatherService)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"temp_C":19,"temp_F":66.2,"temp_K":292.15}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
