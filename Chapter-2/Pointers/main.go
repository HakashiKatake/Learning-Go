package main

import "fmt"

func main() {
	var p *int

	fmt.Println(p) // <nil>

	a := 10

	var pq *int = &a

	fmt.Println("address:", pq) //adress
}
