package promotion_test

import (
	"github.com/isavita/shop/product"
	"github.com/isavita/shop/promotion"
	"testing"
)

func TestGeneralPromotionApply(t *testing.T) {
	gp := promotion.GeneralPromotion{ThresholdAmount: 50}

	if gp.Apply(100) != true {
		t.Error("expected promotion to apply")
	}

	if gp.Apply(10) != false {
		t.Error("expected promotion to not apply")
	}
}

func TestGeneralPromotionDiscount(t *testing.T) {
	gp := promotion.GeneralPromotion{ThresholdAmount: 50, DiscountPercentage: 25}

	if gp.Discount(100) != 25 {
		t.Error("expected promotion discount to be 25")
	}

	if gp.Discount(10) != 0 {
		t.Error("expected promotion discount to be 0")
	}
}

func TestProductPromotionApply(t *testing.T) {
	products := product.Products{
		product.Product{Code: "001", Name: "Kids T-shirt", Price: 1000},
		product.Product{Code: "002", Name: "Cufflinks", Price: 3000},
	}

	pp := promotion.ProductPromotion{ProductCode: "001", MinProductCount: 2}

	if pp.Apply(products) != false {
		t.Error("expected promotion to not apply")
	}

	products = append(products, product.Product{Code: "001", Name: "Kids T-shirt", Price: 1000})

	if pp.Apply(products) != true {
		t.Error("expected promotion to apply")
	}
}

func TestProductPromotionDiscount(t *testing.T) {
	products := product.Products{
		product.Product{Code: "001", Name: "Kids T-shirt", Price: 1000},
		product.Product{Code: "002", Name: "Cufflinks", Price: 3000},
	}

	pp := promotion.ProductPromotion{ProductCode: "001", MinProductCount: 2, DiscountPercentage: 10}

	if pp.Discount(products) != 0 {
		t.Error("expected promotion discount to be 0")
	}

	products = append(products, product.Product{Code: "001", Name: "Kids T-shirt", Price: 1000})

	if pp.Discount(products) != 200 {
		t.Error("expected promotion discount to be 200")
	}
}
