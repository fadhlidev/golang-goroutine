package main

import (
	"fmt"
	"sync"
	"testing"
)

func CreateSingleton() {
	fmt.Println("Create Singleton")
}

func TestOne(t *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			// This will only run once!
			once.Do(CreateSingleton)
			wg.Done()
		}()
	}

	wg.Wait()
}
