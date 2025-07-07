package repository

import (
	"errors"
	"github.com/marcosfrancomarinho/go-first-project/src/application/mappers"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
	"gorm.io/gorm"
)

type GormFindorUser struct {
}

func (g *GormFindorUser) FindByEmail(user *entities.UserLogin) (*entities.UserRegister, error) {
	var client database.Database
	var User database.User
	var userLoginMappers mappers.UserLoginMappers

	client.Connection()

	result := client.DB.Where("Email = ?", user.GetEmail()).First(&User)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("usuario não foi encontrado")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userFound, err := userLoginMappers.GetUser(&User)
	if err != nil {
		return nil, err
	}

	return userFound, nil
}
