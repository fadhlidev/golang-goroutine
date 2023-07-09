package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	var wg sync.WaitGroup

	// Create a new pool
	pool := sync.Pool{
		// Called as a fallback when `Get` function returns <nil>
		New: func() interface{} {
			return "LERO LERO LERO"
		},
	}

	// Add data to the pool
	pool.Put("ORA")
	pool.Put("MUDA")

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Get data from the pool
			data := pool.Get()
			fmt.Println(data)

			// Put the data back to the pool
			pool.Put(data)
		}()
	}

	wg.Wait()
}
