package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func AsyncProcess(wg *sync.WaitGroup, id int) {
	// Decrement delta from the wait group
	defer wg.Done()

	fmt.Printf("Start goroutine : #%d\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finish goroutine : #%d\n", id)
}

func TestWaitGroup(t *testing.T) {
	// Create a new wait group
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		// Add delta to the wait group
		wg.Add(1)
		go AsyncProcess(wg, i)
	}

	// Waits until the wait group's delta becomes zero
	wg.Wait()
	fmt.Println("Finish all goroutines")
}
