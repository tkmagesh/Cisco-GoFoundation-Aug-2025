package main

import "fmt"

// CAN have unused variables at package scope
var x int

// CANNOT use := at package scope
// x := 100

func main() {
	/*
		var userName string
		userName = "Magesh"
	*/

	// var userName string = "Magesh"

	// type inference (works only when declaration & initialization in the same statement)
	// var userName = "Magesh"

	// Using := (can be used only in a function scope)
	userName := "Magesh"
	fmt.Printf("Hello %s, Have a nice day!\n", userName)

	// cannot have unused variables in function scope
	// var x int
}
