package valueobject

import (
	"github.com/google/uuid"
	"time"
)

//Transaction is payment between two parties
type Transaction struct {
	// all value lowercase since they are immutable
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}

// I set all the entities as pointers,
// this is because an entity can change state
// ,and I want that to reflect across all instances of the runtime that has access to it.
// The value objects are held as nonPointer though since they cannot change state.
