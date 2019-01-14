package product_test

import (
	"github.com/isavita/shop/product"
	"testing"
)

func TestProduct(t *testing.T) {
	product := product.Product{"001", "Kids T-shirt", 999}

	if product.Code != "001" {
		t.Error("expected code: \"001\"")
	}

	if product.Name != "Kids T-shirt" {
		t.Error("expected name: \"Kids T-shirt\"")
	}

	if product.Price != 999 {
		t.Error("expected price: 999")
	}
}

func TestProducts(t *testing.T) {
	products := product.Products{
		product.Product{"001", "Kids T-shirt", 999},
		product.Product{"002", "Cufflinks", 1000},
	}

	if len(products) != 2 {
		t.Error("expected 2 products")
	}
}
