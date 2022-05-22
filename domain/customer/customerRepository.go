// Package Customer holds all the domain logic for the customer domain

package customer

import (
	"ddd-go/aggregate"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
)

// CustomerRepository is  interface that defines the rules around what a customer repository
// Has to be able to perform
type CustomerRepository interface {
	Get(uuid uuid.UUID) (aggregate.Customer, error)
	Add(customer aggregate.Customer) error
	Update(customer aggregate.Customer) error
}
