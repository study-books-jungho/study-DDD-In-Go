package entity

import "github.com/google/uuid"

// Item represents Item for all sub domain
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
