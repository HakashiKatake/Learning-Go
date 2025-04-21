package main

import "fmt"

func main() {
	//var m map[K]V
	var m map[string]int

	fmt.Println(m)

	//built-in function
	// make(map[K]V, cap)

	var m1 = make(map[string]int)

	fmt.Println(m1)

	//map literal

	var m3 = map[string]int{"a": 0, "b": 1}

	fmt.Println(m3)

	//map with struct
	TypeFuc()

	var m4 = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}

	fmt.Println(m4)

	MapAdd()

	//retrive value

	var m5 = map[string]int{"a": 0, "b": 1}
	fmt.Println(m5["a"])
	fmt.Println(m5["b"])

	//exists
	Exists()
	//update
	Update()
	//delete
	Delete()
	//iteration
	Iteration()
	//properties
	// 1. Maps are reference types
	// 2. Maps are unordered
	// 3. Maps are not safe for concurrent use
	// 4. Maps are not comparable
	// 5. Maps are not safe for concurrent use

	Properties()

}

type User struct {
	Name string
}

func TypeFuc() {
	var m = map[string]User{
		"a": User{"Peter"},
		"b": User{"Seth"},
	}

	fmt.Println(m)
}

func MapAdd() {
	var m = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}

	m["c"] = User{"Steve"}

	fmt.Println(m)

	var m1 = map[string]int{"a": 0, "b": 1}

	m1["c"] = 2
	fmt.Println(m1)

}

func Exists() {

	var m = map[string]int{"a": 0, "b": 1}
	fmt.Println(m)
	c, ok := m["c"]
	fmt.Println("Key c:", c, ok)

	d, ok := m["d"]
	fmt.Println("Key d:", d, ok)

}

func Update() {
	var m = map[string]int{"a": 0, "b": 1}
	fmt.Println(m)

	m["a"] = 2
	fmt.Println(m)

	m["b"] = 3
	fmt.Println(m)
}

func Delete() {
	var m = map[string]int{"a": 0, "b": 1}
	fmt.Println(m)

	delete(m, "a")
	fmt.Println(m)

	delete(m, "b")
	fmt.Println(m)
}

func Iteration() {
	var m = map[string]int{"a": 0, "b": 1}
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	for k := range m {
		fmt.Printf("Key: %s\n", k)
	}

	for _, v := range m {
		fmt.Printf("Value: %d\n", v)
	}

}

func Properties() {
	type User struct {
		Name string
	}

	var m1 = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}

	m2 := m1
	m2["c"] = User{"Steve"}

	fmt.Println(m1) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
	fmt.Println(m2) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
}
