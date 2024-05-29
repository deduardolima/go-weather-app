package main

import (
	"log"
	"net/http"

	"github.com/deduardolima/go-weather-app/internal/location"
	"github.com/deduardolima/go-weather-app/internal/weather"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := viper.GetString("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatalf("WEATHER_API_KEY not set")
	} else {
		log.Printf("WEATHER_API_KEY loaded: %s", apiKey)
	}

	locationService := location.NewLocationService()
	weatherService := weather.NewWeatherClient()

	r := mux.NewRouter()
	weatherHandler := weather.NewWeatherHandler(locationService, weatherService)
	r.Handle("/weather", weatherHandler).Methods("GET")

	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
