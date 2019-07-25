package main

import (
	"testing"
)

func TestGetsWeather(t * testing.T) {
	cur, lat := GetWeather()

	// if len(result.Summaries) != 2 {
	// 	t.Errorf("Not the right number of summaries!")
	// }

	if cur.Summary == "" {
		t.Errorf("Did not set current summary!")
	}

	if lat.Summary == "" {
		t.Errorf("Did not set next summary!")
	}	
}
