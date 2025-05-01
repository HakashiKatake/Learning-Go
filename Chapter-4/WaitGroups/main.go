package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Task", id, "is running")
}

func main() {
	// WaitGroups
	// WaitGroups are used to wait for a collection of goroutines to finish.
	// They are part of the sync package in Go.
	// A WaitGroup is used to wait for a collection of goroutines to finish.
	// It is used to synchronize the completion of multiple goroutines.
	// Example:
	// var wg sync.WaitGroup
	// wg.Add(1) // Add 1 to the WaitGroup counter
	// go func() {
	// 	defer wg.Done() // Decrement the counter when the goroutine completes
	// }()
	// wg.Wait() // Block until the counter is zero

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Add 1 to the WaitGroup counter
		go task(i, &wg)

	}

	wg.Wait() // Block until the counter is zero
}
