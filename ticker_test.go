package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	// Ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

func TestTick(t *testing.T) {
	// Ticks every 1 second, but it returns the channel immediately
	tick := time.Tick(1 * time.Second)

	for tm := range tick {
		fmt.Println(tm)
	}
}
