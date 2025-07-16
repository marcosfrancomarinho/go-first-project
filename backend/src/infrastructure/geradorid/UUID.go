package geradorid

import (
	"github.com/google/uuid"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type UUID struct {
}

func NewUUID() interfaces.GeratorID {
	return &UUID{}
}

func (u *UUID) Generate() (*valuesobject.ID, error) {
	hash := uuid.New().String()
	id, err := valuesobject.NewID(hash)
	return id, err
}
