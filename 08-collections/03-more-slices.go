package main

import "fmt"

func main() {
	// var nos []int
	nos := make([]int, 0, 3)
	// nos := make([]int, 3, 3)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 2)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 1)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 4)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 3)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 5)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 9)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 7)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos2 := nos[2:4]
	fmt.Printf("len(nos2) : %d, cap(nos2) : %d, nos2 : %v\n", len(nos2), cap(nos2), nos2)
}
