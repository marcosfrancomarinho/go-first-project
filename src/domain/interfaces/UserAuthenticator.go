package interfaces

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type UserAuthenticator interface {
	GenerateToken(user *entities.UserRegister) (*valuesobject.Token, error)
	ValidateToken(token *valuesobject.Token) error
}
