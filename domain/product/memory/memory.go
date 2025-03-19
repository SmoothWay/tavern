package memory

import (
	"sync"

	"github.com/SmoothWay/tavern/domain/product"

	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mr *MemoryProductRepository) GetAll() ([]product.Product, error) {
	products := make([]product.Product, 0)
	for _, product := range mr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mr *MemoryProductRepository) GetById(id uuid.UUID) (product.Product, error) {
	mr.Lock()
	p, ok := mr.products[id]
	if !ok {
		return p, product.ErrProductNotFount
	}
	mr.Unlock()
	return p, nil
}
func (mr *MemoryProductRepository) Add(p product.Product) error {
	mr.Lock()
	if _, ok := mr.products[p.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}
	mr.products[p.GetID()] = p
	mr.Unlock()
	return nil
}
func (mr *MemoryProductRepository) Update(p product.Product) error {
	mr.Lock()
	if _, ok := mr.products[p.GetID()]; !ok {
		return product.ErrProductNotFount
	}
	mr.products[p.GetID()] = p
	mr.Unlock()
	return nil
}

func (mr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mr.Lock()
	if _, ok := mr.products[id]; !ok {
		return product.ErrProductNotFount
	}
	delete(mr.products, id)
	mr.Unlock()
	return nil
}
