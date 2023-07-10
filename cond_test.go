package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func WaitCondition(cond *sync.Cond, wg *sync.WaitGroup, value int) {
	// Lock the condition
	cond.L.Lock()

	// Ask the condition if it can run to the next line
	// needs to call `cond.Signal()` or `cond.Broadcast()` somewhere to run / continue to the next line
	cond.Wait()
	fmt.Printf("Value : %d\n", value)
	wg.Done()

	// Unlock the condition
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	// Create a new Condition using Mutex as the locker
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WaitCondition(cond, &wg, i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)

			// Send a signal to the condition to allow a waiting goroutine to continue run
			cond.Signal()

			// Send a broadcast to the condition to allow all waiting goroutines to continue run
			//cond.Broadcast()
		}
	}()

	wg.Wait()
}
