package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i < 2 {
			continue
		}

		fmt.Println(i)

		if i > 5 {
			break
		}
	}

	fmt.Println("We broke out!")

	//we can omit the init and post statements

	i := 0

	for i < 10 {
		i += 1
	}

	//forever loop

	for {
		fmt.Println("Hello")

	}
}
