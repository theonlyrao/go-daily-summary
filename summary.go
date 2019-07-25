package main

import (
	"fmt"
	"sync"
)

func main() {
        var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		tweet, time := GetLatestTweet()
		fmt.Println("Pres. Trump says: " + tweet)
		fmt.Println(time)
		fmt.Println()
		wg.Done()
	} ()

	wg.Add(1)
	go func() {
		cur, lat := GetWeather()
		fmt.Println(cur.Summary)
		fmt.Println(lat.Summary)
		fmt.Println()
		wg.Done()
	} ()

	wg.Wait()
}
