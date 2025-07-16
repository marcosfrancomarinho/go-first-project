package mappers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
)

type UserLoginMappers struct {
}

func NewUserLoginMappers() *UserLoginMappers {
	return &UserLoginMappers{}
}

func (u *UserLoginMappers) GetUser(user *database.User) (*entities.UserRegister, error) {
	name, err := valuesobject.NewName(user.Name)
	if err != nil {
		return nil, err
	}

	email, err := valuesobject.NewEmail(user.Email)
	if err != nil {
		return nil, err
	}

	password, err := valuesobject.NewPassword(user.Password)
	if err != nil {
		return nil, err
	}

	id, err := valuesobject.NewID(user.Id)
	if err != nil {
		return nil, err
	}

	userFound := entities.NewUserRegister(name, email, password, id)

	return userFound, nil
}
