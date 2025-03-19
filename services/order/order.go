package order

import (
	"context"
	"log"

	"github.com/SmoothWay/tavern/domain/customer"
	"github.com/SmoothWay/tavern/domain/customer/memory"
	"github.com/SmoothWay/tavern/domain/customer/mongo"
	prodmem "github.com/SmoothWay/tavern/domain/product/memory"

	"github.com/SmoothWay/tavern/domain/product"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository
}

func New(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerId uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	// Fetch the customer

	c, err := o.customers.Get(customerId)
	if err != nil {
		return 0, err
	}
	var products []product.Product
	var total float64

	for _, id := range productsIDs {
		product, err := o.products.GetById(id)
		if err != nil {
			return 0, err
		}
		products = append(products, product)
		total += product.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products with total price: %f", c.GetID(), len(products), total)
	return total, nil
}

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	cust, err := customer.New(name)
	if err != nil {
		return uuid.Nil, err
	}

	err = o.customers.Add(cust)
	if err != nil {
		return uuid.Nil, err
	}

	return cust.GetID(), nil
}
