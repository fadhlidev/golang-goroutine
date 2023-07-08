package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// Create a new channel
	c := make(chan string)
	// :sparkles: It is recommended to always close the channel
	defer close(c)

	go func() {
		fmt.Println("Sending data")
		time.Sleep(3 * time.Second)

		// Send data to the channel
		c <- "Hello there"
	}()

	// Receive data from the channel
	data := <-c
	fmt.Printf("Receive data: %s\n", data)
}

// Channel can only send data
func SendMessage(c chan<- string, message string) {
	fmt.Printf("Sending message: %s\n", message)
	time.Sleep(2 * time.Second)

	c <- message
}

// Channel can only receive data
func ReceiveMessage(c <-chan string) {
	message := <-c

	fmt.Printf("Recieve message: %s\n", message)
}

func TestInOutChannel(t *testing.T) {
	c := make(chan string)
	defer close(c)

	go ReceiveMessage(c)
	go SendMessage(c, "Hello there")

	time.Sleep(3 * time.Second)
}
