package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeatherResponse struct {
	Properties WeatherProperties
}

type WeatherProperties struct {
	Periods []WeatherPeriods
}

type WeatherPeriods struct {
	Name string
	DetailedForecast string
}

type WeatherSummary struct {
	Summary string
}

func GetWeather() (current * WeatherSummary, later * WeatherSummary) {
	client := &http.Client {
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest("GET", "https://api.weather.gov/gridpoints/EWX/166,85/forecast", nil)
	req.Header.Add("User-Agent", "A-script")
	
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("error: " + err.Error())
		return nil, nil
	}

	defer resp.Body.Close()

	weather := new(WeatherResponse)
	json.NewDecoder(resp.Body).Decode(weather)	

	current = new(WeatherSummary)
	later = new(WeatherSummary)
	
	current.Summary = weather.Properties.Periods[0].Name + ": " + weather.Properties.Periods[0].DetailedForecast;
	later.Summary = weather.Properties.Periods[1].Name + ": " + weather.Properties.Periods[1].DetailedForecast

	return current, later
}
