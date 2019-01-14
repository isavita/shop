package product

import (
	"fmt"
)

type Product struct {
	Code  string
	Name  string
	Price int
}

func (p Product) ProductInfo() {
	fmt.Println("Product:")
	fmt.Printf("\tcode: %s\n", p.Code)
	fmt.Printf("\tname: %s\n", p.Name)
	fmt.Printf("\tprice: %.2f\n", float64(p.Price)/100.0)
}

type Products []Product
