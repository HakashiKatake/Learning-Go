package main

import "fmt"

func main() {
	//switch case
	// switch case statement is used to execute a block of code based on a value. It is similar to if else statement, but it is more readable and easier to maintain.

	day := "monday"

	switch day {
	case "monday":
		fmt.Println("time to work!")
	case "friday":
		fmt.Println("let's party")
	default:
		fmt.Println("browse memes")
	}

	//shorthand switch

	switch day := "monday"; day {
	case "monday":
		fmt.Println("time to work!")
	case "friday":
		fmt.Println("let's party")
	default:
		fmt.Println("browse memes")
	}

	//fallthrough

	switch day := "monday"; day {
	case "monday":
		fmt.Println("time to work!")
		fallthrough // this keyword can transfer control to the next case
	case "friday":
		fmt.Println("let's party")
	default:
		fmt.Println("browse memes")

	}

	//conditional

	x := 10

	switch {
	case x > 5:
		fmt.Println("x is greater")
	default:
		fmt.Println("x is not greater")
	}

}
