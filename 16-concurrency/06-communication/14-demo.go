package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	/*
		data := <-ch
		ch <- 100
		fmt.Println(data)
	*/

	/*
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	/*
		go func() {
			ch <- 100 //2.(NB)
		}()
		data := <-ch // 1.(B) 3.(UB)
		fmt.Println(data)
	*/

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		data := <-ch // 2.(NB)
		fmt.Println(data)
	}()
	ch <- 100 // 1.(B) 3.(UB)
	wg.Wait()
}
