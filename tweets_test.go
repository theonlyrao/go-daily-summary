package main

import (
	"testing"
)

func TestGetsTrumpsLatestTweet(t * testing.T) {
	tweet, _ := GetLatestTweet()

	if tweet == "" {
		t.Errorf("Retrieved tweet is empty")
	}
}
