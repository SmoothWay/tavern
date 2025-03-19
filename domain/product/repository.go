package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFount      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("such product already exists")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id uuid.UUID) (Product, error)
	Add(product Product) error
	Update(prodcut Product) error
	Delete(id uuid.UUID) error
}
