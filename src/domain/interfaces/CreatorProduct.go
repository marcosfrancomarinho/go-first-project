package interfaces

import "github.com/marcosfrancomarinho/go-first-project/src/domain/entities"

type CreatorProduct interface {
	Create(product *entities.Product) error
}
