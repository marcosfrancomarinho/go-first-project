package interfaces

import e "github.com/marcosfrancomarinho/go-first-project/src/domain/entities"

type CreateUser interface {
	Create(user *e.UserRegister) error
}
