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

func GetWeather() (* http.Response) {
	client := &http.Client {
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest("GET", "https://api.weather.gov/gridpoints/EWX/166,85/forecast", nil)
	req.Header.Add("User-Agent", "A-script")
	
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("error: " + err.Error())
		return nil
	}

	defer resp.Body.Close()

	weather := new(WeatherResponse)
	json.NewDecoder(resp.Body).Decode(weather)	

	fmt.Println(weather.Properties.Periods[0].Name + ": " + weather.Properties.Periods[0].DetailedForecast)
	fmt.Println(weather.Properties.Periods[1].Name + ": " + weather.Properties.Periods[1].DetailedForecast)	
	
	return resp
}
