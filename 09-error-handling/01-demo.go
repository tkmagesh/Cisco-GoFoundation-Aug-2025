package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	divisor := 7
	// divisor := 0
	quotient, reminder, err := divide(100, divisor)
	if err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d and reminder = %d\n", divisor, quotient, reminder)
		return
	} else if err == ErrDivideByZero {
		fmt.Println("Please do not attempt to divide by 0")
	} else {
		fmt.Println("Unknown Error :", err)
	}

}

func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	fmt.Println("[divide] Calculating quotient")
	if divisor == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = multiplier / divisor
	fmt.Println("[divide] Calculating remainder")
	remainder = multiplier % divisor
	return
}
