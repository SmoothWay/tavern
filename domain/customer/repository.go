package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("customer not found in repository")
	ErrFailedToCreateCustomer = errors.New("failed to create customer")
	ErrUpdateCustomer         = errors.New("failed to update customer")
	ErrCustomerAlreadyExists  = errors.New("customer already exists")
)

type CustomerRepository interface {
	Get(id uuid.UUID) (Customer, error)
	Add(Customer) error
	Update(Customer) error
}
