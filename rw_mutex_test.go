package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Account struct {
	mutex sync.RWMutex
	point int
}

// Lock when mutating data
func (a *Account) AddPoint(p int) {
	a.mutex.Lock()
	a.point += p
	a.mutex.Unlock()
}

// Lock when getting data
func (a *Account) getPoint() int {
	a.mutex.RLock()
	p := a.point
	a.mutex.RUnlock()
	return p
}

func TestReadWriteMutex(t *testing.T) {
	account := Account{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddPoint(1)
				fmt.Printf("Current point : %d\n", account.getPoint())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Final point : %d\n", account.getPoint())
}
