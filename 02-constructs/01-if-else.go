package main

import "fmt"

func main() {
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	// scope of the "no" variable is limited to the if-else block
	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	// "no" is not accessible outside if-else block
	// fmt.Println("no =", no)
}
