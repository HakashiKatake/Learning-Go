package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var c int = 30
	var d int = 40

	fmt.Println("Arithmetic Operators")
	fmt.Println("Addition: ", a+b)
	fmt.Println("Subtraction: ", a-b)
	fmt.Println("Multiplication: ", a*b)
	fmt.Println("Division: ", a/b)
	fmt.Println("Modulus: ", a%b)

	fmt.Println("Relational Operators")
	fmt.Println("Less than: ", a < b)
	fmt.Println("Greater than: ", a > b)
	fmt.Println("Less than or equal to: ", a <= b)
	fmt.Println("Greater than or equal to: ", a >= b)
	fmt.Println("Equal to: ", a == b)
	fmt.Println("Not equal to: ", a != b)

	fmt.Println("Logical Operators")
	fmt.Println("Logical AND: ", (a < b) && (c < d))
	fmt.Println("Logical OR: ", (a < b) || (c > d))
	fmt.Println("Logical NOT: ", !(a < b))

	// Bitwise Operators
	a = 5      // 0101
	b = 3      // 0011
	c = a & b  // 0001
	d = a | b  // 0111
	e := a ^ b // 0110

	fmt.Println("Bitwise Operators")
	fmt.Printf("%d & %d = %d\n", a, b, c)
	fmt.Printf("%d | %d = %d\n", a, b, d)
	fmt.Printf("%d ^ %d = %d\n", a, b, e)

	// Assignment Operators
	a += 5 // Same as a = a + 5
	b -= 3 // Same as b = b - 3
	c *= 2 // Same as c = c * 2
	d /= 4 // Same as d = d / 4

	fmt.Println("Assignment Operators")
	fmt.Printf("a += 5 : %d\n", a)
	fmt.Printf("b -= 3 : %d\n", b)
	fmt.Printf("c *= 2 : %d\n", c)
	fmt.Printf("d /= 4 : %d\n", d)

	// Unary Operators
	a++ // Increment operator
	b-- // Decrement operator

	fmt.Println("Unary Operators")
	fmt.Printf("a++ : %d\n", a)
	fmt.Printf("b-- : %d\n",
		b)
	fmt.Printf("++a : %d\n", a)
	fmt.Printf("--b : %d\n", b)

	// Ternary Operator
	// Go does not have a ternary operator like C/C++. Instead, you can use an if-else statement.

	//complex number

	var c1 complex128 = complex(10, 1)
	var c2 complex64 = 12 + 4i

	// complex128 is a complex number with 128-bit precision
	// complex64 is a complex number with 64-bit precision
	// complex numbers are represented as a + bi, where a is the real part and b is the imaginary part

	fmt.Println("Complex Number")
	fmt.Println("Complex number 1: ", c1)
	fmt.Println("Complex number 2: ", c2)

}
