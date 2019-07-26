package main

import (
	"fmt"
)

func showLatestTweet(tweet, time chan string) {
	w, i := GetLatestTweet()
	tweet <- w
	time <- i
}

func showWeather(current, later chan string) {
	c, l := GetWeather()
	current <- c.Summary
	later <- l.Summary
}

func main() {
	current := make(chan string)
	later := make(chan string)

	tweet := make(chan string)
	time := make(chan string)

	go showLatestTweet(tweet, time)
	go showWeather(current, later)		
	
	for x := 0; x < 4; x++ {
		select {
		case w := <- tweet:
			fmt.Println(w)
		case i := <- time:
			fmt.Println(i)
		case c := <- current:
			fmt.Println(c)
		case l := <- later:
			fmt.Println(l)
		}
	}
}
