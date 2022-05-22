// file : product.go
// product is an aggregate that represents a product

package aggregate

import (
	"ddd-go/entity"
	"errors"
	"github.com/google/uuid"
)

var (
	//ErrMissingValues is returned when a product is created without a name or description
	ErrMissingValues = errors.New("missing values")
)

//Product is a aggregate that combine item with a price and quantity
type Product struct {
	// item is the root entity which is an item
	item  *entity.Item
	price float64
	// Quantity is the number of products in stock
	quantity int
}

// NewProduct will create a new product
// will return error if name of description is empty
func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}
func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
