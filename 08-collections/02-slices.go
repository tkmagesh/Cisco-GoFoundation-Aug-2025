package main

import "fmt"

func main() {
	var nos []int
	nos = append(nos, 3)
	nos = append(nos, 2, 1)

	// var x []int = []int{5,4}
	// var x = []int{5,4}
	x := []int{5, 4}

	nos = append(nos, x...)
	fmt.Println(nos)

	fmt.Println("Iterating an slice - using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an slice - using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	nos2 := nos // a new reference is created
	nos2[0] = 9999
	fmt.Printf("nos2[0] = %d, nos = %d\n", nos2[0], nos[0])

	fmt.Println("Before sorting :", nos)
	sort(nos)
	fmt.Println("After sorting :", nos)
}

func sort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
