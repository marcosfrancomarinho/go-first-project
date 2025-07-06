package interfaces

import e "github.com/marcosfrancomarinho/src/domain/entities"

type FindorUser interface {
	FindByEmail(user *e.UserLogin) *e.UserRegister
}
