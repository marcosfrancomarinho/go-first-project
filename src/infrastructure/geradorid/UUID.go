package geradorid

import (
	"github.com/google/uuid"
	V "github.com/marcosfrancomarinho/src/domain/valuesobject"
)

type UUID struct {
}

func (u *UUID) Generate() (*V.ID, error) {
	hash := uuid.New().String()
	id, err := V.NewID(hash)
	return id, err
}
