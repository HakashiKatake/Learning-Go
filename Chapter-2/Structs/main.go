package main

import "fmt"

func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	//muliple fields
	type Person2 struct {
		FirstName, LastName string
		Age                 int
	}

	var p1 Person

	fmt.Println("Person 1:", p1)

	//struct literal
	var p2 = Person2{FirstName: "Karan", LastName: "Pratap Singh", Age: 22}

	fmt.Println("Person 2:", p2)

	fmt.Println("Person 2 First Name:", p2.FirstName)
	fmt.Println("Person 2 Last Name:", p2.LastName)

	ptr := &p1

	fmt.Println((*ptr).FirstName)

	Embedding()
	composition()
}

func Embedding() {

	type Person struct {
		FirstName, LastName string
		Age                 int
	}

	type SuperHero struct {
		Person
		Alias string
	}

	s := SuperHero{}

	s.FirstName = "Bruce"
	s.LastName = "Wayne"
	s.Age = 40
	s.Alias = "batman"

	fmt.Println(s)

}

func composition() {
	type Person struct {
		FirstName, LastName string
		Age                 int
	}

	type SuperHero struct {
		Person Person
		Alias  string
	}

	p := Person{"Bruce", "Wayne", 40}
	s := SuperHero{p, "batman"}

	fmt.Println(s)
}
