package repository

import (
	"errors"
	"github.com/marcosfrancomarinho/go-first-project/src/application/mappers"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
	"gorm.io/gorm"
)

type GormFindorUser struct {
}

func NewGormFindorUser() interfaces.FindorUser {
	return &GormFindorUser{}
}

func (g *GormFindorUser) FindByEmail(userLogin *entities.UserLogin) (*entities.UserRegister, error) {
	user := &database.User{}
	userLoginMappers := mappers.NewUserLoginMappers()

	client := database.NewDataBase()
	result := client.Where("Email = ?", userLogin.GetEmail()).First(user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("usuario não foi encontrado")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userFound, err := userLoginMappers.GetUser(user)
	if err != nil {
		return nil, err
	}

	return userFound, nil
}
