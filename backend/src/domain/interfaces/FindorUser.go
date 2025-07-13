package interfaces

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
)

type FindorUser interface {
	FindByEmail(user *entities.UserLogin) (*entities.UserRegister, error)
}
