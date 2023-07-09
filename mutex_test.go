package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounterWithoutMutex(t *testing.T) {
	counter := 0

	// This will simulate a race condition
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				counter = counter + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Total count: %d\n", counter)
}

func TestCounterWithMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				counter = counter + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Total count: %d\n", counter)
}
