package main

import "fmt"

func main() {

	//simple function declaration

	//func myFunc() {}

	//function calling
	//myFunc()

	myFunction("Hello")
	myNextFunction("Hello", "World")
	s := Function("Hello")
	fmt.Println(s)
	s1, s2 := Function2("Hello")
	fmt.Println(s1, s2)
	s3, s4 := Function3("Hello")
	fmt.Println(s3, s4)

}

//cant make function inside a function

func myFunction(p1 string) {
	fmt.Println(p1)
}

// short hand function declaration
func myNextFunction(p1, p2 string) {}

// function with return type after () it is the return type of the function
func Function(p1 string) string {
	msg := fmt.Sprintf("%s function", p1)
	return msg
}

// function with multiple return types
func Function2(p1 string) (string, string) {
	msg := fmt.Sprintf("%s function", p1)
	return msg, "Hello"
}

// function with named return types

func Function3(p1 string) (msg string, msg2 string) {
	msg = fmt.Sprintf("%s function", p1)
	msg2 = "Hello"
	return //this is known as naked return

}
