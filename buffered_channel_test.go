package main

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	// Create a new channel with 5 capacities data (Channel w/ buffer)
	c := make(chan string, 3)
	defer close(c)

	c <- "Jonathan"
	c <- "Joseph"
	c <- "Jotaro"

	// It doesn't matter if there is no goroutine tries to get any data from the channel
}

func TestBufferedChannelWithReceiver(t *testing.T) {
	c := make(chan string, 3)
	defer close(c)

	c <- "Jonathan"
	c <- "Joseph"
	c <- "Jotaro"

	// No matter how fast the sender is, the program will slow if the receiver is slow
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("Hello, %s!\n", <-c)
		}
	}()

	time.Sleep(4 * time.Second)
}
