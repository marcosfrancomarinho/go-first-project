package interfaces

import e "github.com/marcosfrancomarinho/src/domain/entities"

type CreateUser interface {
	Create(user *e.UserRegister) error
}
