package main

import (
	"fmt"
	"github.com/isavita/shop/checkout"
	"github.com/isavita/shop/product"
	"github.com/isavita/shop/promotion"
)

func printBasket(basket *product.Products) {
	fmt.Printf("\nBasket: ")
	limit := len(*basket) - 1
	for i, product := range *basket {
		if i < limit {
			fmt.Printf("%s,", product.Code)
		} else {
			fmt.Printf("%s", product.Code)
		}
	}
}

func main() {
	product1 := product.Product{Code: "001", Name: "Lavender heart", Price: 925}
	product2 := product.Product{Code: "002", Name: "Personalised cufflinks", Price: 4500}
	product3 := product.Product{Code: "003", Name: "Kids T-shirt", Price: 1995}

	fmt.Println("Test data")
	fmt.Println("---------")
	product1.ProductInfo()
	product2.ProductInfo()
	product3.ProductInfo()

	generalPromotions := []promotion.GeneralPromotion{
		promotion.GeneralPromotion{ThresholdAmount: 6000, DiscountPercentage: 10},
	}

	productPromotions := []promotion.ProductPromotion{
		promotion.ProductPromotion{ProductCode: "001", MinProductCount: 2, DiscountAmount: 75},
	}

	checkout1 := checkout.Checkout{ProductPromotions: productPromotions, GeneralPromotions: generalPromotions}
	checkout1.Scan(product1)
	checkout1.Scan(product2)
	checkout1.Scan(product3)
	printBasket(&checkout1.Products)
	fmt.Printf("\nTotal price expected: %.2f\n", float64(checkout1.Total())/100.0)

	checkout2 := checkout.Checkout{ProductPromotions: productPromotions, GeneralPromotions: generalPromotions}
	checkout2.Scan(product1)
	checkout2.Scan(product3)
	checkout2.Scan(product1)
	printBasket(&checkout2.Products)
	fmt.Printf("\nTotal price expected: %.2f\n", float64(checkout2.Total())/100.0)

	checkout3 := checkout.Checkout{ProductPromotions: productPromotions, GeneralPromotions: generalPromotions}
	checkout3.Scan(product1)
	checkout3.Scan(product2)
	checkout3.Scan(product1)
	checkout3.Scan(product3)
	printBasket(&checkout3.Products)
	fmt.Printf("\nTotal price expected: %.2f\n", float64(checkout3.Total())/100.0)
}
