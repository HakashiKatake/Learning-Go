package main

import "fmt"

func init() {
	fmt.Println("Before main!")
}

func main() {
	fmt.Println("Running main")
}

func init() {
	fmt.Println("Before main!")
}

func init() {
	fmt.Println("Hello again?")
}
