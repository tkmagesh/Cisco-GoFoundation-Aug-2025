package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the execution of f1 through the scheduler to be executed in future
	f2()

	// Poor man's synchronization techniques (DO NOT use them)
	// time.Sleep(1 * time.Second)
	// fmt.Scanln()
	time.Sleep(2 * time.Second)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
