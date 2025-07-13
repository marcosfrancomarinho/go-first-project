package interfaces

import "github.com/marcosfrancomarinho/go-first-project/src/domain/entities"

type FinderProduct interface {
	FindAll() (*[]entities.Product, error)
}
