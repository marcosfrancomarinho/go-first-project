package interfaces

import "github.com/marcosfrancomarinho/go-first-project/src/domain/entities"

type FindorProduct interface {
	FindAll() (*[]entities.Product, error)
}
