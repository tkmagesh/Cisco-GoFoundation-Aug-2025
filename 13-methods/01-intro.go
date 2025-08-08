package main

import "fmt"

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

func main() {
	/*
		p1 := Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	p1 := &Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	/*
		fmt.Println(Format(p1))
		ApplyDiscount(&p1, 10)
		fmt.Println(Format(p1))
	*/

	fmt.Println(p1.Format())
	p1.ApplyDiscount(10)
	fmt.Println(p1.Format())
}
