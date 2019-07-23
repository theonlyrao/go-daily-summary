package main

import (
	"testing"
)

func TestGetsTrumpsLatestTweet(t * testing.T) {
	result := GetLatestTweet()

	if len(result) == 0 {
		t.Errorf("Retrieved tweet is empty")
	}
}
