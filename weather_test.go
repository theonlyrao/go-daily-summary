package main

import (
	"testing"
)

func TestGetsWeather(t * testing.T) {
	result := GetWeather()

	if len(result.Summaries) != 2 {
		t.Errorf("Not the right number of summaries!")
	}

	if result.Summaries[0].Summary == "" {
		t.Errorf("Did not set current summary!")
	}

	if result.Summaries[1].Summary == "" {
		t.Errorf("Did not set next summary!")
	}	
}
