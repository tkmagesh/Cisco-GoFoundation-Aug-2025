package main

import "fmt"

func main() {
	exec(f1) // invoke "exec()" in such a way that - f1 is invoked
	exec(f2) // invoke "exec()" in such a way that - f2 is invoked
	exec(func() {
		fmt.Println("anon invoked")
	})
}

func exec(fn func()) {
	/* execute either f1 or f2 based on the argument */
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
