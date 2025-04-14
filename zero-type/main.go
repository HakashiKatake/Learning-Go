package main

import "fmt"

func main() {
	// Zero value types in Go
	// In Go, every variable has a default value, known as the zero value.
	// The zero value is the value that a variable has when it is declared but not initialized.
	// The zero value is different for each type:
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

}
