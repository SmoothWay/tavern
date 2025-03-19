// package aggregate holds our aggrets that combines many entities into a full object

package customer

import (
	"errors"

	"github.com/SmoothWay/tavern"

	"github.com/google/uuid"
)

var (
	ErrinvalidPerson = errors.New("customer should have valid name")
)

type Customer struct {
	// person is the root entity
	// which means person.ID is the main identifier for the customer
	person   *tavern.Person
	products []*tavern.Item

	transaction []tavern.Transaction
}

// NewCustomer is a factory to create a new customer aggregate
// it will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrinvalidPerson
	}
	person := &tavern.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:      person,
		products:    make([]*tavern.Item, 0),
		transaction: make([]tavern.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.ID = uuid.New()
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
