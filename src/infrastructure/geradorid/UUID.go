package geradorid

import (
	"github.com/google/uuid"
	"github.com/marcosfrancomarinho/src/domain/valuesobject"
)

type UUID struct {
}

func (u *UUID) Generate() (*valuesobject.ID, error) {
	hash := uuid.New().String()
	id, err := valuesobject.NewID(hash)
	return id, err
}
