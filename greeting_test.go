package main

import (
	"fmt"
	"testing"
	"time"
)

func Greet(name string) {
	fmt.Printf("(approaching %s)\n", name)
	time.Sleep(2 * time.Second)
	fmt.Printf("Hello %s!\n", name)
}

func TestGreetSequential(t *testing.T) {
	Greet("Dio")
	Greet("Jonathan")
}

func TestGreetAsynchronous(t *testing.T) {
	// Run a new go routine using `go` keyword followed by a function
	// The function will be executed asynchronously
	go Greet("Dio")
	go Greet("Jonathan")

	time.Sleep(3 * time.Second)
}
