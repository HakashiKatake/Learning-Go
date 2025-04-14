package main

import "fmt"

func main() {
	//if else
	// if else statement is used to execute a block of code based on a condition. statement doesnt have to be inside (), its optional

	x := 10

	if x > 5 {
		fmt.Println("x is gt 5")
	} else if x > 10 {
		fmt.Println("x is gt 10")
	} else {
		fmt.Println("else case")
	}

	//compact if

	if x := 10; x > 5 {
		fmt.Println("x is gt 5")
	}

}
