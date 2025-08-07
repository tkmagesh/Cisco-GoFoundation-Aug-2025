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
	// divisor := 0
	var divisor int
	for {
		fmt.Println("Enter the divisor :")

		// DONT validate the divisor

		fmt.Scanln(&divisor)
		if quotient, reminder, err := divideAdapter(100, divisor); err == ErrDivideByZero {
			fmt.Println("Enter a non zero divisor")
			continue
		} else if err != nil {
			fmt.Println("Unknown error :", err)
			break
		} else {
			fmt.Printf("Dividing 100 by %d, quotient = %d and reminder = %d\n", divisor, quotient, reminder)
			break
		}
	}

}

func divideAdapter(multiplier, divisor int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = ErrDivideByZero
		}
	}()
	quotient, remainder = divide(multiplier, divisor)
	return
}

// 3rd party library api (can't change the code)
func divide(multiplier, divisor int) (quotient, remainder int) {
	fmt.Println("[divide] Calculating quotient")
	quotient = multiplier / divisor
	fmt.Println("[divide] Calculating remainder")
	remainder = multiplier % divisor
	return
}
