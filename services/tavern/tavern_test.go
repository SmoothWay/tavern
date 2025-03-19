package tavern

import (
	"context"
	"testing"

	"github.com/SmoothWay/tavern/domain/product"
	"github.com/SmoothWay/tavern/services/order"
	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := order.New(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	custId, err := os.AddCustomer("Bob")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(custId, order)
	if err != nil {
		t.Fatal(err)
	}
}

func init_products(t *testing.T) []product.Product {
	beer, err := product.New("German beer", "Made in Germany", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	meat, err := product.New("Steak", "500g beef", 10.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.New("Red wine", `healthy wine 10%% of alcohol`, 10)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{
		beer, meat, wine,
	}
}
