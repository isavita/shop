package checkout_test

import (
	"github.com/isavita/shop/checkout"
	"github.com/isavita/shop/product"
	"github.com/isavita/shop/promotion"
	"testing"
)

func TestCheckoutScan(t *testing.T) {
	checkout := checkout.Checkout{}

	if len(checkout.Products) != 0 {
		t.Errorf("expected to contain zero products, but it has %v", len(checkout.Products))
	}

	checkout.Scan(product.Product{})

	if len(checkout.Products) != 1 {
		t.Errorf("expected to contain one products, but it has %v", len(checkout.Products))
	}
}

func TestCheckoutTotal(t *testing.T) {
	product1 := product.Product{"001", "Lavender heart", 925}
	product2 := product.Product{"002", "Personalised cufflinks", 4500}
	product3 := product.Product{"003", "Kids T-shirt", 1995}

	generalPromotions := []promotion.GeneralPromotion{
		promotion.GeneralPromotion{ThresholdAmount: 6000, DiscountPercentage: 10},
	}

	productPromotions := []promotion.ProductPromotion{
		promotion.ProductPromotion{ProductCode: "001", MinProductCount: 2, DiscountPercentage: 10},
	}

	checkout := checkout.Checkout{ProductPromotions: productPromotions, GeneralPromotions: generalPromotions}
	checkout.Scan(product1)
	checkout.Scan(product2)
	checkout.Scan(product3)

	if checkout.Total() != 6678 {
		t.Errorf("expected total amount after discount to be 6678, but it is %v", checkout.Total())
	}

	checkout.Scan(product1)

	if checkout.Total() != 7343 {
		t.Errorf("expected total amount after discount to be 7343, but it is %v", checkout.Total())
	}
}
