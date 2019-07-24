package main

import (
	"fmt"
	"io/ioutil"	
	"net/http"
	"time"
)

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
	body, _ := ioutil.ReadAll(resp.Body)
	weather := string(body)
	fmt.Println(weather)

	return resp
}
