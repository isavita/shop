package main

import (
	"fmt"
	"github.com/isavita/shop/product"
)

func main() {
	var product1 product.Product
	product1 = product.Product{Code: "001", Name: "Lavender heart", Price: 925}
	product2 := product.Product{"002", "Personalised cufflinks", 4500}
	product3 := &product.Product{"003", "Kids T-shirt", 1995}

	fmt.Println("Test data")
	product1.ProductInfo()
	product2.ProductInfo()
	product3.ProductInfo()

	basket1 := product.Products{product1, product2, *product3}

	fmt.Printf("Basket: ")
	limit := len(basket1) - 1
	for i, product := range basket1 {
		if i < limit {
			fmt.Printf("%s, ", product.Code)
		} else {
			fmt.Printf("%s", product.Code)
		}
	}
	fmt.Printf("\nTotal price expected: %.2f", 5.6)
}
