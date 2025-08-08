package main

import (
	"fmt"
	"time"
)

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Dummy struct {
	CreatedAt time.Time
}

// PerishableProduct = Product + Expiry

type PerishableProduct struct {
	Product // a new attribute 'Product' is auto created
	Expiry  string
	Dummy
}

// factory function
func NewPerishableProduct(id int, name string, cost float64, createdAt time.Time, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Dummy: Dummy{
			CreatedAt: createdAt,
		},
		Expiry: expiry,
	}
}

func main() {

	/*
		milk := PerishableProduct{
			Product: Product{
				Id:   101,
				Name: "Milk",
				Cost: 50,
			},
			Dummy: Dummy{
				CreatedAt: time.Now(),
			},
			Expiry: "2 Days",
		}
	*/

	milk := NewPerishableProduct(101, "Milk", 50, time.Now(), "2 Days")
	fmt.Printf("%#v\n", milk)
	fmt.Println("milk.Product.Id =", milk.Product.Id)
	fmt.Println("milk.Id =", milk.Id)
	fmt.Println("milk.CreatedAt =", milk.CreatedAt)

	// Try using the below 'Format' and 'ApplyDiscount' functions for 'milk'
	fmt.Println(Format(milk.Product))
	ApplyDiscount(&milk.Product, 10)
	fmt.Println(Format(milk.Product))
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
