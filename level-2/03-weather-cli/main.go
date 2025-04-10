package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherData struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		TempF     float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Humidity    int     `json:"humidity"`
		WindKph     float64 `json:"wind_kph"`
		FeelsLikeC  float64 `json:"feelslike_c"`
		FeelsLikeF  float64 `json:"feelslike_f"`
		UV         float64 `json:"uv"`
	} `json:"current"`
}

func getWeather(apiKey, location string) (*WeatherData, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, location)
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var weather WeatherData
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("error parsing weather data: %v", err)
	}

	return &weather, nil
}

func displayWeather(weather *WeatherData) {
