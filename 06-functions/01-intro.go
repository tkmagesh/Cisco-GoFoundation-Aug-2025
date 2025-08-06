package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHi()
	greet("Magesh")
	greetUser("Magesh", "Kuppan")
	fmt.Println(getGreetMsg("Suresh", "Kannan"))

	/*
		firstName, lastName := getNames("Magesh Kuppan")
		fmt.Printf("FirstName : %s, LastName : %s\n", firstName, lastName)
	*/

	firstName, _ := getNames("Magesh Kuppan")
	fmt.Printf("FirstName : %s\n", firstName)

}

// No arguments, No return results
func sayHi() {
	fmt.Println("Hi!")
}

// 1 argument, No return results
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

// 2 arguments, No return results
/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}
*/

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}

// 2 arguments, 1 result
func getGreetMsg(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a wonderful day!", firstName, lastName)
}

// 2 arguments, 2 results
/*
func getNames(fullName string) (string, string) {
	var firstName, lastName string
	names := strings.Split(fullName, " ")
	firstName = names[0]
	lastName = names[1]
	return firstName, lastName
}
*/

// Using named results
func getNames(fullName string) (firstName, lastName string) {
	names := strings.Split(fullName, " ")
	/*
		firstName = names[0]
		lastName = names[1]
	*/
	firstName, lastName = names[0], names[1]
	return
}
