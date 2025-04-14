package main

import "fmt"

type MyAlias = string

type MyDefined string

func main() {
	var alias MyAlias
	//var def MyDefined

	// ✅ Works
	var copy1 string = alias

	// ❌ Cannot use def (variable of type MyDefined) as string value in variable
	//var copy2 string = def

	fmt.Println(copy1)
}
