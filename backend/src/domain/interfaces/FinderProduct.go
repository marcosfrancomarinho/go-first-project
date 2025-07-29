package interfaces

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type FindorProduct interface {
	FindAll(container *entities.ContainerPage) (*[]entities.Product, error)
	TotalCount() (*valuesobject.Quantity, error)
}
