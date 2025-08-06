package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var fn func()
	for range 10 {
		fn = getFn()
		fn()
		time.Sleep(500 * time.Millisecond)
	}
}

func getFn() func() {
	if rand.Intn(20)%2 == 0 {
		return f1
	}
	return f2
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
