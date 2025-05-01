package main

import (
	"fmt"
	"time"
)

//  func task(id int) {
// 	fmt.Println("Task", id, "is running")
// }

func main() {
	// Goroutines
	// Goroutines are lightweight threads managed by the Go runtime.
	// They are used to perform concurrent tasks in a Go program.
	// A goroutine is created using the `go` keyword followed by a function call.
	// The function will run concurrently with the rest of the program.
	// Example:
	// go func() {
	// 	fmt.Println("Hello from a goroutine!")
	// }()

	for i := 1; i <= 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2) // Wait for goroutines to finish

}
