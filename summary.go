package main

import (
	"fmt"
	"sync"
)

func main() {
        var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		GetLatestTweet()
		fmt.Println()
		wg.Done()
	} ()

	wg.Add(1)
	go func() {
		weather := GetWeather()
		fmt.Println(weather.Summaries[0].Summary)
		fmt.Println(weather.Summaries[1].Summary)
		fmt.Println()
		wg.Done()
	} ()

	wg.Wait()
}
