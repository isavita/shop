package product_test

import (
	"github.com/isavita/shop/product"
	"testing"
)

func TestProduct(t *testing.T) {
	product := product.Product{"001", "Kids T-shirt", 999}

	if product.Code != "001" {
		t.Error("expected \"001\"")
	}

	if product.Name != "Kids T-shirt" {
		t.Error("expected \"Kids T-shirt\"")
	}

	if product.Price != 999 {
		t.Error("expected 999")
	}
}
