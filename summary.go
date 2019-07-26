package main

import (
	"fmt"
	"sync"
)

func showLatestTweet(twitter chan string) {
	w, i := GetLatestTweet()
	twitter <- w
	twitter <- i
	close(twitter)
}

func showWeather(weather chan string) {
	c, l := GetWeather()
	weather <- c.Summary
	weather <- l.Summary
	close(weather)
}

func merge(cs ...<-chan string) <-chan string {
    var wg sync.WaitGroup
    out := make(chan string)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan string) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func main() {
	weather := make(chan string)
	twitter := make(chan string)

	go showWeather(twitter)
	go showLatestTweet(weather)
	
	for output := range merge(weather, twitter) {
		fmt.Println(output)
	}
}
