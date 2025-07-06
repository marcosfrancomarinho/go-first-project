package mappers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
)

type UserLoginMappers struct {
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
	userFound, err := entities.NewUserRegister(name, email, password, id)
	if err != nil {
		return nil, err
	}
	return userFound, nil
}
