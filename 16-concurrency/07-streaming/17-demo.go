package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go generateNos(ch)
LOOP:
	for {
		if no, isOpen := <-ch; isOpen {
			fmt.Println(no)
			continue LOOP
		}
		break LOOP
	}
	/*
		for no := range ch {
			fmt.Println(no)
		}
	*/
	fmt.Println("Done!")
}

// producer
func generateNos(ch chan int) {
	count := rand.Intn(20)
	fmt.Printf("[generateNos] about to produce %d numbers\n", count)
	for i := range count {
		time.Sleep(500 * time.Millisecond)
		ch <- (i + 1) * 10
	}
	close(ch)
}
