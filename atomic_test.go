package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	var counter int64

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// Atomic: allows mutating primitive values w/o without using locker (e.g. Mutex)
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Total counter : %d\n", counter)
}
