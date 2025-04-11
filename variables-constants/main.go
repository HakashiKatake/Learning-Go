package main

import "fmt"

var outOfScope string = "This is out of scope" //this is a global variable, it can be used in any function

func main() {
	var inScope string = "This is in scope" //this is a local variable, it can only be used in this function
	var name string = "Saurabh"
	var hello, world string = "Hello", "World" //this is a multiple variable declaration
	var yo = "Yo"                              //this is a variable declaration with type inference
	foo := "Foo"                               //this is a short variable declaration, it can only be used in this function

	fmt.Println("Hi there " + name)
	fmt.Println("This is a local variable: " + inScope)
	fmt.Println("This is a global variable: " + outOfScope)
	fmt.Println("This is a multiple variable declaration: " + hello + " " + world)
	fmt.Println("This is a variable declaration with type inference: " + yo)
	fmt.Println("This is a short variable declaration: " + foo)

	constant()
}

func constant() {
	const pi = "3.14" //this is a constant, it cannot be changed

	fmt.Println("This is a constant: " + pi)
}
