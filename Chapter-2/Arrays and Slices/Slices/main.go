package main

import "fmt"

func main() {
	a := [5]int{20, 15, 5, 30, 25}

	s := a[1:4]

	// Output: Array: [20 15 5 30 25], Length: 5, Capacity: 5
	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", a, len(a), cap(a))

	// Output: Slice [15 5 30], Length: 3, Capacity: 4
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", s, len(s), cap(s))

	//declaring a slice

	// var s1 []int

	var s2 []string

	fmt.Println(s2)
	fmt.Println(s2 == nil)

	//bult-in function
	// make([]T, len, cap)

	var s3 = make([]int, 0, 0)

	fmt.Println(s3)

	//low and high
	// s3[low:high]

	var b = [4]string{
		"C++",
		"Go",
		"Java",
		"TypeScript",
	}

	ss := b[0:2] // Select from 0 to 2
	ss1 := b[:3] // Select first 3
	ss2 := b[2:] // Select from 1 to end

	fmt.Println("Array:", b)
	fmt.Println("Slice 1:", ss)
	fmt.Println("Slice 2:", ss1)
	fmt.Println("Slice 3:", ss2)

	//copying a slice

	//func copy(dst, src []T) int
	Copy()

	//appending a slice
	Append()

	//properties of slices
	// 1. Slices are reference types
	// 2. Slices are dynamic
	// 3. Slices are flexible
	// 4. Slices are built on top of arrays

	Property()
}

func Copy() {
	s1 := []string{"a", "b", "c", "d"}
	s2 := make([]string, len(s1))

	e := copy(s2, s1)

	fmt.Println("Src:", s1)
	fmt.Println("Dst:", s2)
	fmt.Println("Elements:", e)
}

func Append() {
	//func append(slice []T, elements ...T) []T

	s1 := []string{"a", "b", "c", "d"}

	s2 := append(s1, "e", "f")

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}

func Property() {
	a := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	s := a[0:2]

	s[0] = "Sun"

	fmt.Println(a) // Output: [Sun Tue Wed Thu Fri Sat Sun]
	fmt.Println(s) // Output: [Sun Tue]
}
