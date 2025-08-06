package main

import (
	"fmt"
	"strings"
)

func main() {
	// No arguments, No return results
	func() {
		fmt.Println("Hi!")
	}()

	// 1 argument, No return results
	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	// 2 arguments, No return results
	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	// 2 arguments, 1 result
	msg := func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a wonderful day!", firstName, lastName)
	}("Suresh", "Kannan")
	fmt.Println(msg)

	// 2 arguments, 2 results (named results)
	firstName, lastName := func(fullName string) (firstName, lastName string) {
		names := strings.Split(fullName, " ")
		firstName, lastName = names[0], names[1]
		return
	}("Magesh Kuppan")
	fmt.Printf("FirstName : %s, LastName : %s\n", firstName, lastName)

}
