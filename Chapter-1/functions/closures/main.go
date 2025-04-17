package main

import "fmt"

func main() {
	addd := myFunction()

	addd(5)
	fmt.Println(addd(10))

	sum := add(1, 2, 3, 5)
	fmt.Println(sum)
}

func myFunction() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v

		return sum
	}
}

//variadic function

func add(values ...int) int {
	sum := 0

	for _, v := range values {
		sum += v
	}

	return sum
}
