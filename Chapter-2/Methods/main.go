package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Car struct {
	Name string
	Year int
}

func (c Car) IsLatest() bool {
	return c.Year >= 2017
}

//	func (c Car) UpdateName(name string) {
//		c.Name = name
//	}
//
// pointer type
func (c *Car) UpdateName(name string) {
	c.Name = name
}

type MyInt int

func (i MyInt) isGreater(value int) bool {
	return i > MyInt(value)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	c := Car{"Tesla", 2021}

	fmt.Println("IsLatest", c.IsLatest())

	c.UpdateName("Tesla Model S")
	fmt.Println("Updated Name", c.Name)

	i := MyInt(10)

	fmt.Println(i.isGreater(5))

}
