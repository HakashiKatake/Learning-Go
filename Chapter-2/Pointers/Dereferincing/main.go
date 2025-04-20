package main

import "fmt"

func main() {
	a := 10

	var p *int = &a

	fmt.Println("address:", p)
	fmt.Println("value:", *p)

	change()
	newWay()
	p2p()
}

func change() {
	a := 10

	var p *int = &a

	fmt.Println("before", a)
	fmt.Println("address:", p)

	*p = 20
	fmt.Println("after:", a)
}

//pointer func

//myFunction(&a)
//func myFunction(ptr *int) {}

func newWay() {
	p := new(int)
	*p = 100

	fmt.Println("value", *p)
	fmt.Println("address", p)
}

//pointer to pointer

func p2p() {
	p := new(int)
	*p = 100

	p1 := &p

	fmt.Println("P value", *p, " address", p)
	fmt.Println("P1 value", *p1, " address", p)

	fmt.Println("Dereferenced value", **p1)
}
