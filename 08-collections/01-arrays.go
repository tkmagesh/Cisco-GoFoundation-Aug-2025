package main

import "fmt"

func main() {
	// memory for 5 ints are allocated and initialized with the 'zero value' of int
	// var nos [5]int

	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an array - using indexer")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an array - using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		nos2 := nos // a copy is created
		nos2[0] = 9999
		fmt.Println("nos[0] = ", nos[0])
	*/
	var nos2 *[5]int
	nos2 = &nos
	fmt.Println("(*nos2)[0] :", (*nos2)[0])

	// accessing member of a "container" pointer does not require dereferencing
	fmt.Println("nos2[0] :", nos2[0])

	fmt.Println("Before sorting :", nos)
	sort(&nos)
	fmt.Println("After sorting :", nos)
}

func sort(list *[5]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
