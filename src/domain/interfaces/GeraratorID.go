package interfaces

import v "github.com/marcosfrancomarinho/src/domain/valuesobject"

type GeratorID interface {
	Generate() (*v.ID, error)
}
