package entity

import "github.com/google/uuid"

//Person is entity that represents a person in all domain
type Person struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID uuid.UUID

	// Name is the name of the person
	Name string

	// Age is the age of the person
	Age int
}
