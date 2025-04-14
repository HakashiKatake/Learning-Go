package main

import "fmt"

func main() {
	var name string = "Saurabh"
	var bio string = `
		Hello, my name is Saurabh.
		I am a software engineer.
		I love to code in Go.`

	var value bool = true
	var isItTrue bool = false
	var age int = 25
	//numeric types

	var i int = 404                     // Platform dependent
	var i8 int8 = 127                   // -128 to 127
	var i16 int16 = 32767               // -2^15 to 2^15 - 1
	var i32 int32 = -2147483647         // -2^31 to 2^31 - 1
	var i64 int64 = 9223372036854775807 // -2^63 to 2^63 - 1

	fmt.Println("This is a string: " + name)
	fmt.Println("This is a multi-line string: " + bio)
	fmt.Println("This is a boolean: " + fmt.Sprintf("%t", value))
	fmt.Println("This is a boolean: " + fmt.Sprintf("%t", isItTrue))
	fmt.Println("This is an integer: " + fmt.Sprintf("%d", age))
	fmt.Println("This is a numeric type")
	fmt.Println("Signed integer: ")
	fmt.Println("This is an integer: " + fmt.Sprintf("%d", i))
	fmt.Println("This is an integer of 8bit: " + fmt.Sprintf("%d", i8))
	fmt.Println("This is an integer of 16bit: " + fmt.Sprintf("%d", i16))
	fmt.Println("This is an integer of 32bit: " + fmt.Sprintf("%d", i32))
	fmt.Println("This is an integer of 64bit: " + fmt.Sprintf("%d", i64))

	//unsigned types
	var ui uint = 404                     // Platform dependent
	var ui8 uint8 = 255                   // 0 to 255
	var ui16 uint16 = 65535               // 0 to 2^16
	var ui32 uint32 = 2147483647          // 0 to 2^32
	var ui64 uint64 = 9223372036854775807 // 0 to 2^64
	var uiptr uintptr                     // Integer representation of a memory address

	fmt.Println("Unsigned integer: ")
	fmt.Println("This is an unsigned integer: " + fmt.Sprintf("%d", ui))
	fmt.Println("This is an unsigned integer of 8bit: " + fmt.Sprintf("%d", ui8))
	fmt.Println("This is an unsigned integer of 16bit: " + fmt.Sprintf("%d", ui16))
	fmt.Println("This is an unsigned integer of 32bit: " + fmt.Sprintf("%d", ui32))
	fmt.Println("This is an unsigned integer of 64bit: " + fmt.Sprintf("%d", ui64))
	fmt.Println("This is an unsigned integer of pointer: " + fmt.Sprintf("%d", uiptr))

	//  Byte and rune
	type byte = uint8
	type rune = int32

	var b byte = 255        //Golang has two additional integer types called byte and rune that are aliases for uint8 and int32 data types respectively.
	var r rune = 1234567890 //A rune represents a unicode code point.
	fmt.Println("Byte and rune: ")

	fmt.Println("This is a byte: " + fmt.Sprintf("%d", b))
	fmt.Println("This is a rune: " + fmt.Sprintf("%d", r))

	//  Float types

	//go has two floating point types: float32 and float64

	var f32 float32 = 1.7812 // IEEE-754 32-bit
	var f64 float64 = 3.1415 // IEEE-754 64-bit

	fmt.Println("Floating point types: ")
	fmt.Println("This is a float 32: " + fmt.Sprintf("%f", f32))
	fmt.Println("This is a float 64: " + fmt.Sprintf("%f", f64))

}
