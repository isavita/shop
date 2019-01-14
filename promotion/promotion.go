package promotion

import (
	"github.com/isavita/shop/product"
	"math"
)

type GeneralPromoter interface {
	Apply(amount int) bool
	Discount(amount int) int
}

type GeneralPromotion struct {
	ThresholdAmount    int
	DiscountPercentage int
}

func (gp GeneralPromotion) Apply(amount int) bool {
	if gp.ThresholdAmount > amount {
		return false
	}

	return true
}

func (gp GeneralPromotion) Discount(amount int) int {
	if gp.Apply(amount) {
		return int(math.Round((float64(gp.DiscountPercentage) / 100.0) * float64(amount)))
	}

	return 0
}

type ProductPromoter interface {
	Apply(products product.Products) bool
	Discount(products product.Products) int
}

type ProductPromotion struct {
	ProductCode        string
	MinProductCount    int
	DiscountPercentage int
}

func (pp ProductPromotion) Apply(products product.Products) bool {
	matches := 0
	for _, product := range products {
		if pp.ProductCode == product.Code {
			matches += 1
			if pp.MinProductCount == matches {
				return true
			}
		}
	}

	return false
}

func (pp ProductPromotion) Discount(products product.Products) int {
	var discount int
	if pp.Apply(products) {
		for _, product := range products {
			if pp.ProductCode == product.Code {
				discount += int(math.Round((float64(pp.DiscountPercentage) / 100.0) * float64(product.Price)))
			}
		}
	}

	return discount
}
