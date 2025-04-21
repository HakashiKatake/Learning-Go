package main

import "fmt"

// var a [n]T
func main() {
	var arr [4]int

	fmt.Println(arr)

	//var a [n]T = [n]T{V1, V2, ... Vn}

	var arr1 [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//short hand

	arr2 := [4]int{1, 2, 3, 4}
	fmt.Println(arr2)

	//acessing array elements
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])

	//iteration

	for i := 0; i < len(arr2); i++ {
		fmt.Println("Index: %d, Value: %d", i, arr2[i])
	}

	//for range

	for i, e := range arr {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}

	// for i, e := range arr {} // Normal usage of range

	// for _, e := range arr {} // Omit index with _ and use element

	// for i := range arr {} // Use index only

	// for range arr {} // Simply loop over the array

	// multi-dimensional array

	arr3 := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, e := range arr3 {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}

	//size inference

	arr4 := [...]int{1, 2, 3, 4}
	fmt.Println(arr4)

	//copying arrays

	var a = [4]int{1, 2, 3, 4}
	var b = a // copy of a

	b[0] = 10      // changing b does not affect a
	fmt.Println(a) // [1 2 3 4]
	fmt.Println(b) // [10 2 3 4]

}
