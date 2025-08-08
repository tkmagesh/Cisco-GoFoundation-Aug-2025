package main

import "fmt"

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Dolor sint nisi commodo velit minim tempor adipisicing non tempor voluptate id nisi culpa anim."
	x = true
	x = 99.99
	fmt.Println(x)

	// x = 100
	x = "Ullamco dolor elit excepteur fugiat mollit nisi laboris commodo amet."
	// y := x * 2

	// unsafe
	// y := x.(int) * 2

	// safe (assert the type before you take an action)
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// type assertion using type-switch
	// x = 100
	// x = "Dolor sint nisi commodo velit minim tempor adipisicing non tempor voluptate id nisi culpa anim."
	// x = true
	// x = 99.99
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a strinng, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, x * 0.01 =", val*0.01)
	default:
		fmt.Println("x is of unknown type")
	}

}
