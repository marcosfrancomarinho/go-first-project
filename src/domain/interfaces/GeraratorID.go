package interfaces

import v "github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"

type GeratorID interface {
	Generate() (*v.ID, error)
}
