package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("Application panicked with error :", err)
			return
		}
		fmt.Println("Thank you!")
	}()

	// divisor := 7
	divisor := 0
	quotient, reminder := divide(100, divisor)
	fmt.Printf("Dividing 100 by %d, quotient = %d and reminder = %d\n", divisor, quotient, reminder)

}

func divide(multiplier, divisor int) (quotient, remainder int) {
	fmt.Println("[divide] Calculating quotient")
	quotient = multiplier / divisor
	fmt.Println("[divide] Calculating remainder")
	remainder = multiplier % divisor
	return
}
