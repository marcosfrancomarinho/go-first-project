package interfaces

import e "github.com/marcosfrancomarinho/go-first-project/src/domain/entities"

type FindorUser interface {
	FindByEmail(user *e.UserLogin) *e.UserRegister
}
