package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	panic("dummy panic")
	wg.Add(10)
	go f1() // scheduling the execution of f1 through the scheduler to be executed in future
	f2()

	wg.Wait() // block until the counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
