package main

import (
	"testing"
)

func TestGetsWeather(t * testing.T) {
	result := GetWeather()

	if result == nil {
		t.Errorf("Retrieved weather is empty")
	}
}
