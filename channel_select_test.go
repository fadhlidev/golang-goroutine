package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)
	defer func() {
		close(c1)
		close(c2)
	}()

	for i := 0; i < 5; i++ {
		go func(i int) {
			c1 <- fmt.Sprintf("C1#%d", i)
		}(i)

		go func(i int) {
			c1 <- fmt.Sprintf("C2#%d", i)
		}(i)
	}

	// Receives data from two channels using `select` keyword
	counter := 0
	for counter < 10 {
		select {
		case data := <-c1:
			fmt.Printf("Data from channel 1 : %s\n", data)
			counter++
		case data := <-c2:
			fmt.Printf("Data from channel 2 : %s\n", data)
			counter++
		}
	}
}

func TestChannelDefaultSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)
	defer func() {
		close(c1)
		close(c2)
	}()

	for i := 0; i < 5; i++ {
		go func(i int) {
			c1 <- fmt.Sprintf("C1#%d", i)
			time.Sleep(1 * time.Second)
		}(i)

		go func(i int) {
			c1 <- fmt.Sprintf("C2#%d", i)
			time.Sleep(1 * time.Second)
		}(i)
	}

	counter := 0
	for counter < 10 {
		select {
		case data := <-c1:
			fmt.Printf("Data from channel 1 : %s\n", data)
			counter++
		case data := <-c2:
			fmt.Printf("Data from channel 2 : %s\n", data)
			counter++
		default:
			fmt.Println("Waiting to receive data...")
		}
	}
}
