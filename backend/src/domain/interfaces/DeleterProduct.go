package interfaces

import "github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"

type DeleterProduct interface {
	Delete(id *valuesobject.ID) error
}
