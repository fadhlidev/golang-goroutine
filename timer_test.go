package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Printf("Current time : %s\n", time.Now().String())

	// Receive time from the timer after its duration is expired
	// This will block the code
	tm := <-timer.C
	fmt.Printf("Timer time : %s\n", tm.String())
}

func TestAfter(t *testing.T) {
	// Returns the channel immediately after duration expired
	c := time.After(5 * time.Second)

	tick := <-c
	fmt.Println(tick)
}

func TestAfterFunc(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	// AfterFunc will run asynchronously and execute callback function after duration expired
	time.AfterFunc(5*time.Second, func() {
		wg.Done()
		fmt.Println("Runs after 5 seconds")
	})

	wg.Wait()
}
