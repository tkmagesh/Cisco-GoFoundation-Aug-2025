package main

import (
	"fmt"
	"strings"
)

func main() {
	// No arguments, No return results
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()

	// 1 argument, No return results
	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	// 2 arguments, No return results
	var greetUser func(string, string)
	greetUser = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}
	greetUser("Magesh", "Kuppan")

	// 2 arguments, 1 result
	var getGreetMsg func(string, string) string
	getGreetMsg = func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a wonderful day!", firstName, lastName)
	}
	msg := getGreetMsg("Suresh", "Kannan")
	fmt.Println(msg)

	// 2 arguments, 2 results (named results)
	var getNames func(string) (string, string)
	getNames = func(fullName string) (firstName, lastName string) {
		names := strings.Split(fullName, " ")
		firstName, lastName = names[0], names[1]
		return
	}
	firstName, lastName := getNames("Magesh Kuppan")
	fmt.Printf("FirstName : %s, LastName : %s\n", firstName, lastName)

}
