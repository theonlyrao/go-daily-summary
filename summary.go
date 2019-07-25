package main

import (
	"fmt"
)

func main() {
	GetLatestTweet()
	fmt.Println()	

	weather := GetWeather()
	fmt.Println(weather.Summaries[0].Summary)
	fmt.Println(weather.Summaries[1].Summary)
}
