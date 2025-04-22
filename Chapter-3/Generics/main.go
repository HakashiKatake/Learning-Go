package main

import "fmt"

// func fnName[T constraint]() {
// 	...
// }

type SumConstraint interface {
	int | float64 | string
}

func summ[T SumConstraint](a, b T) T {
	return a + b
}

func main() {
	sum(1, 2)
	sum(4.0, 2.0)
	sum("a", "b")
	summ(1, 2)
	fmt.Println(summ(4.0, 2.0))
}

type Addable interface {
	~int | ~float64 | ~string
}

func sum[T Addable](a, b T) T {
	return a + b
}
