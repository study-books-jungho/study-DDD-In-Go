// Package aggregate holds aggregates that combine many entity into a full object

package aggregate

import (
	"ddd-go/entity"
	"ddd-go/valueobject"
	"errors"
	"github.com/google/uuid"
)

var (
	//ErrInvalidPerson is return when  the  person is not valid in the new Customer factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

//Customer is a aggregate that combines all entities needed to represent a  customer
type Customer struct {

	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	person *entity.Person

	// a customer can hold many products
	products []*entity.Item

	// a customer can perform many transactions
	transaction []valueobject.Transaction
}

//NewCustomer is a factory to  create new Customer aggregate
// It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	// validate that  the name is empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	// Create new person and generate ID
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		person:      person,
		products:    make([]*entity.Item, 0),
		transaction: make([]valueobject.Transaction, 0),
	}, nil
}

func (c Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c Customer) GetName() string {
	return c.person.Name
}
