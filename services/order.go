// package services holds all the services that connects repository into a business flow

package services

import (
	"ddd-go/domain/customer"
	"ddd-go/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

//OrderService is a implementation of the OrderService
type OrderService struct {
	customer customer.CustomerRepository
}

// NewOrderService takes a variable amount of OrderConfiguration functions and
// and return a new OrderService Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the orderService
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

//WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that parent functions can take in all the needed parameters
	return func(os *OrderService) error {
		os.customer = cr
		return nil
	}
}

//WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameter, such as connection strings they could be inputted here
	cr := memory.New()
	return WithCustomerRepository(cr)
}

// CreateOrder will chain together all repositories to create a order for a customer
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	// Get the customer
	_, err := o.customer.Get(customerID)
	if err != nil {
		return err
	}

	// Get each Product, Ouch, We need a ProductRepository
	return nil
}
