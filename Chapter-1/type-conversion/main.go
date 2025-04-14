package main

import "fmt"

func main() {

	// Type conversion in Go
	// In Go, you can convert a value from one type to another using the type name as a function.
	// For example, to convert an int to a float64, you can use float64(value).
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println("Converted int to float64:", f)
	fmt.Println("Converted float64 to uint:", u)

}
