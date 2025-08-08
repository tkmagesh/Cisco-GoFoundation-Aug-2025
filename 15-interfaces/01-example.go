package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// utility function
/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Invalid argument")
	}
}
*/

/*
	func PrintArea(x interface{}) {
		switch val := x.(type) {
		case interface{ Area() float64 }:
			fmt.Println("Area :", val.Area())
		default:
			fmt.Println("Invalid argument")
		}
	}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

// Introducing Perimeter()
/* 
Circle = 2 * pi * r
Rectangle = 2 * (h + w)
Square = 4 * side
*/

func PrintPerimeter(x ?){
	fmt.Println("Perimeter :", x.Perimeter())
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
	PrintPerimeter(r)

	s := Square{Side: 16}
	PrintArea(s)
	PrintPerimeter(s)
}
