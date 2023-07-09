package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var wg sync.WaitGroup

	// Create a new map
	var data sync.Map

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Store data to the map
			data.Store(fmt.Sprintf("Data#%d", i), i)
		}(i)
	}

	wg.Wait()

	// Loop through the map
	data.Range(func(key, value any) bool {
		fmt.Printf("%s : %d\n", key, value)
		return true
	})
}
