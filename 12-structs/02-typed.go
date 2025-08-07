package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/* product => id, name, cost */
	/*
		var p1 Product
		p1.Id = 100
		p1.Name = "Pen"
		p1.Cost = 10
	*/

	/*
		var p1 Product  = Product {
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	// type inference
	/*
		var p1 = Product {
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	p1 := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// fmt.Println(p1)
	// fmt.Printf("%#v\n", p1)
	fmt.Printf("%+v\n", p1)

	p2 := p1 //A copy is created
	p2.Cost = 20
	fmt.Printf("p2.Cost = %v, p1.Cost = %v\n", p2.Cost, p1.Cost)

	// To create a reference
	var p1Ptr *Product
	p1Ptr = &p1
	p1Ptr.Cost = 20
	fmt.Printf("p1Ptr.Cost = %v, p1.Cost = %v\n", p1Ptr.Cost, p1.Cost)

	/*
		1. Print the product (using Format)
		2. Apply 10% discount (using ApplyDiscount)
		3. Print the product (using Format)
	*/

	fmt.Println(Format(p1))
	ApplyDiscount(&p1, 10)
	fmt.Println(Format(p1))
}

// return a formatted string of the given product
// ex: Id = 100, Name = "Pen", Cost = 10
func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

// Update the cost to apply the given discount(%) on the given product
func ApplyDiscount(pPtr *Product, discountPercentage float64) /* DO NOT RETURN */ {
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}
