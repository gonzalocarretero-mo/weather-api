package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	Description       string `json:"description"`
	CurrentConditions struct {
		Temp float64 `json:"temp"`
	} `json:"currentConditions"`
}

func FetchWeather(apiKey string, city string) (WeatherResponse, error) {
	url := fmt.Sprint("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/", city, "?unitGroup=metric&key=", apiKey, "&contentType=json")
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return WeatherResponse{}, err
	}

	return data, nil
}
