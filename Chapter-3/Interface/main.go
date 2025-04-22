package main

import "fmt"

type PowerDrawer interface {
	Draw(power int)
}

type mobile struct {
	brand string
}

func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

type laptop struct {
	cpu string
}

func (l laptop) Draw(power int) {
	fmt.Printf("%T -> cpu: %s, power: %d\n", l, l.cpu, power)
}

type toaster struct {
	amount int
}

func (t toaster) Draw(power int) {
	fmt.Printf("%T -> amount: %d, power: %d\n", t, t.amount, power)
}

type kettle struct {
	quantity string
}

func (k kettle) Draw(power int) {
	fmt.Printf("%T -> quantity: %s, power: %d\n", k, k.quantity, power)
}

type socket struct{}

func (socket) Plug(device PowerDrawer, power int) {
	device.Draw(power)
}

func main() {

	m := mobile{"Apple"}
	l := laptop{"Intel i9"}
	t := toaster{4}
	k := kettle{"50%"}

	s := socket{}

	s.Plug(m, 10)
	s.Plug(l, 50)
	s.Plug(t, 30)
	s.Plug(k, 25)

	//empty interface

	//var x interface{}

	//type assertion

	var i interface{} = "hello"

	ss := i.(string)
	fmt.Println(ss)

	var i2 interface{} = "10"
	ss, ok := i2.(string)
	fmt.Println(ss, ok)

	//type switch

	var tt interface{}
	tt = "hello"

	switch tt := tt.(type) {
	case string:
		fmt.Printf("string: %s\n", tt)
	case bool:
		fmt.Printf("boolean: %v\n", tt)
	case int:
		fmt.Printf("integer: %d\n", tt)
	default:
		fmt.Printf("unexpected: %T\n", tt)
	}

	//properties

	type MyInterface interface {
		Method()
	}

	var ii MyInterface

	fmt.Println(ii) // <nil>

	//embedding

	type interface1 interface {
		Method1()
	}

	type interface2 interface {
		Method2()
	}

	type interface3 interface {
		interface1
		interface2
	}
}
