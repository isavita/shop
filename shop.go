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

	fmt.Println(product1)
	fmt.Println(product2)
	fmt.Println((*product3).Name)
}
