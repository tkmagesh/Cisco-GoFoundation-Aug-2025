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

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (pPtr *Product) ApplyDiscount(discountPercentage float64) /* DO NOT RETURN */ {
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}

/* ---------- Perishable Product ----------------- */
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

// override the Product.Format() method
func (pp *PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {
	milk := NewPerishableProduct(101, "Milk", 50, time.Now(), "2 Days")

	fmt.Println(milk.Format())
	milk.ApplyDiscount(10)
	fmt.Println(milk.Format())

}
