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
	var db database.Database
	var userModel database.User
	var userLoginMappers mappers.UserLoginMappers

	db.Connection()
	result := db.DB.Where("Email = ?", user.GetEmail()).First(&userModel)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("usuario não foi encontrado")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userFound, err := userLoginMappers.GetUser(&userModel)
	if err != nil {
		return nil, err
	}

	return userFound, nil
}
