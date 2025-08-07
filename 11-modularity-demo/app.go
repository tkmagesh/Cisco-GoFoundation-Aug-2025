package main

import (
	"fmt"

	// "github.com/tkmagesh/cisco-gofoundation-aug-2025/11-modularity-demo/calculator"
	calc "github.com/tkmagesh/cisco-gofoundation-aug-2025/11-modularity-demo/calculator"

	"github.com/fatih/color"
	"github.com/tkmagesh/cisco-gofoundation-aug-2025/11-modularity-demo/calculator/utils"
)

func run() {
	// fmt.Println("Modularity demo app running....")
	color.Yellow("Modularity demo app running....")
	/*
		fmt.Println("Add Result :", calculator.Add(100, 200))
		fmt.Println("Subtract Result :", calculator.Subtract(100, 200))
		fmt.Println("Operation Count :", calculator.OpCount())
	*/

	// using package alias
	// fmt.Println("Add Result :", calc.Add(100, 200))
	color.Red("Add Result : %d\n", calc.Add(100, 200))

	fmt.Println("Subtract Result :", calc.Subtract(100, 200))
	fmt.Println("Operation Count :", calc.OpCount())

	fmt.Println("Is 21 Odd ? :", utils.IsOdd(21))
}
