package main

import "fmt"

func main() {
	var no int
	no = 100
	var addressOfNo *int

	// value -> reference
	addressOfNo = &no
	fmt.Println(addressOfNo)

	// reference -> value (dereferencing)
	fmt.Println(*addressOfNo)

	// in other words
	fmt.Println(no == *(&no))

}
