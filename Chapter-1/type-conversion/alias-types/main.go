package main

import "fmt"

func main() {
	// Alias types in Go
	// An alias type is a new name for an existing type.
	// It is useful when you want to create a new type that has the same underlying representation as an existing type,
	// but you want to give it a different name.
	type MyInt int
	type YourInt int

	var a MyInt = 42
	var b YourInt = 42

	fmt.Println("MyInt:", a)
	fmt.Println("YourInt:", b)

}
