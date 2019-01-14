package checkout

import (
	"github.com/isavita/shop/product"
	"github.com/isavita/shop/promotion"
)

type Checkout struct {
	ProductPromotions []promotion.ProductPromotion
	GeneralPromotions []promotion.GeneralPromotion
	Products          product.Products
}

func (c *Checkout) Scan(product product.Product) {
	c.Products = append(c.Products, product)
}

func (c Checkout) Total() int {
	var discount int
	for _, promotion := range c.ProductPromotions {
		discount += promotion.Discount(c.Products)
	}

	amount := c.amount() - discount

	for _, promotion := range c.GeneralPromotions {
		discount = promotion.Discount(amount)
		amount -= discount
	}

	return amount
}

func (c Checkout) amount() int {
	amount := 0
	for _, product := range c.Products {
		amount += product.Price
	}

	return amount
}
