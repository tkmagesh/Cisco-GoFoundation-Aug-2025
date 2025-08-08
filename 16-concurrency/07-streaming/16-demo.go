package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go generateNos(ch)
	for range 10 {
		fmt.Println(<-ch)
	}
}

func generateNos(ch chan int) {
	for i := range 10 {
		time.Sleep(500 * time.Millisecond)
		ch <- (i + 1) * 10
	}

}
