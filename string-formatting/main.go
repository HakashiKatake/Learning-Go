package main

import "fmt"

func main() {
	// String formatting in Go
	// In Go, you can format strings using the fmt package.
	// The fmt package provides functions for formatting strings, including Printf, Sprintf, and Errorf.

	//print func does no formatting
	fmt.Print("What", "is", "your", "name?")
	fmt.Print("My", "name", "is", "golang")

	//Println func does no formatting but add newline
	fmt.Println("What", "is", "your", "name?")
	fmt.Println("My", "name", "is", "golang")

	// Printf func does formatting
	name := "golang"

	fmt.Println("What is your name?")
	fmt.Printf("My name is %s", name)

	//annotation
	//annotations are used to format the output

	// %s - string
	// %d - int
	// %f - float
	// %t - bool
	// %v - default format
	// %T - type of the value
	// %x - hexadecimal
	// %X - hexadecimal (uppercase)
	// %b - binary
	// %o - octal

	//most used ones are:
	// %s - string
	// %d - int
	// %f - float
	// %t - bool
	// %v - default format

	percent := (7.0 / 9) * 100
	fmt.Printf("%f", percent)

	fpercent := (7.0 / 9) * 100
	fmt.Printf("%.2f %%", fpercent)

	//String formatting with Sprintf

	//Sprint
	// Sprint is used to format a string and return it as a string
	sf := fmt.Sprint("Hello", "world")
	fmt.Println(sf)

	// Sprintf is used to format a string and return it as a string
	s := fmt.Sprintf("hex:%x bin:%b", 10, 10)
	fmt.Println(s)

	//Sprintln
	sp := fmt.Sprintln("Hello", "world")
	fmt.Println(sp)

	//string literals can be used like this
	msg := `
Hello from
multiline
`

	fmt.Println(msg)

}
