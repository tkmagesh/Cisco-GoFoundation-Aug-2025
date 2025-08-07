package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	/*
		logAdd := wrapLogger(add)
		logSubtract := wrapLogger(subtract)

		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		profiledAdd := wrapProfiler(add)
		profiledSubtract := wrapProfiler(subtract)

		profiledAdd(100, 200)
		profiledSubtract(100, 200)
	*/

	// combining profiling & logging
	/*
		logAdd := wrapLogger(add)
		profiledAdd := wrapProfiler(logAdd)
		profiledAdd(100, 200)

		profiledSubtract := wrapProfiler(wrapLogger(subtract))
		profiledSubtract(100, 200)
	*/

	//
	add := wrapProfiler(wrapLogger(add))
	subtract := wrapProfiler(wrapLogger(subtract))

	add(100, 200)
	subtract(100, 200)
}

// handler functions (frozen)
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

// log wrappers
func logAdd(x, y int) {
	log.Println("Invocation started...")
	add(x, y)
	log.Println("Invocation completed...")
}

func logSubtract(x, y int) {
	log.Println("Invocation started...")
	subtract(x, y)
	log.Println("Invocation completed...")
}

// applying 'commonality variability' on the wrappers
func logOperation(op func(int, int), x, y int) {
	log.Println("Invocation started...")
	op(x, y)
	log.Println("Invocation completed...")
}

// retain the same signature of the handler functions
func wrapLogger(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Invocation started...")
		op(x, y)
		log.Println("Invocation completed...")
	}
}

// adding profiler
func wrapProfiler(op func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		timeTaken := time.Since(start)
		log.Printf("Time taken : %v\n", timeTaken)
	}
}
