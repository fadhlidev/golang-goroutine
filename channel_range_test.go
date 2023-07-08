package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelRange(t *testing.T) {
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			c <- fmt.Sprintf("Data : %d", i)
		}

		// Close the channel
		close(c)
	}()

	// Loop through the channel and receive its data as long as the channel is not closed
	for data := range c {
		fmt.Println(data)
	}
}
