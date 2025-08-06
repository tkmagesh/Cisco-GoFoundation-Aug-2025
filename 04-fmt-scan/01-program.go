package main

import "fmt"

func main() {
	var userName string
	fmt.Println("Enter the username :")

	// Scan fn limitation - the data input cannot have any spaces
	fmt.Scanln(&userName)
	fmt.Printf("Hello %s, Have a nice day!\n", userName)

	var n1, n2 int
	fmt.Println("Enter the numbers")
	fmt.Scanln(&n1, &n2)
	fmt.Println(n1, n2)
}
