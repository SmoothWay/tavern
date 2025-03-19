// package memory is a in-memory implementation of CustomerRepository

package memory

import (
	"sync"

	"github.com/SmoothWay/tavern/domain/customer"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	mr.Lock()
	defer mr.Unlock()
	c, ok := mr.customers[id]
	if !ok {
		return c, customer.ErrCustomerNotFound
	}
	return c, nil
}

func (mr *MemoryRepository) Update(req customer.Customer) error {
	mr.Lock()
	defer mr.Unlock()

	_, ok := mr.customers[req.GetID()]
	if !ok {
		return customer.ErrUpdateCustomer
	}
	mr.customers[req.GetID()] = req
	return nil
}

func (mr *MemoryRepository) Add(req customer.Customer) error {
	mr.Lock()
	defer mr.Unlock()

	_, ok := mr.customers[req.GetID()]
	if ok {
		return customer.ErrCustomerAlreadyExists
	}
	mr.customers[req.GetID()] = req
	return nil
}
