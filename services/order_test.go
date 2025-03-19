package services

import (
	"testing"

	"github.com/SmoothWay/tavern/domain/customer"
	"github.com/SmoothWay/tavern/domain/product"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("German beer", "Made in Germany", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	meat, err := product.NewProduct("Steak", "500g beef", 10.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Red wine", `healthy wine 10%% of alcohol`, 10)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{
		beer, meat, wine,
	}
}

func TestOrderNewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Bob")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
