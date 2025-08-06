package main

import "fmt"

func main() {
	var f float64
	var x int8 = 100

	// type conversion syntax = use the typename like a function
	f = float64(x)
	fmt.Println(f)
}
