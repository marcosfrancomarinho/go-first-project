package repository

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
	"github.com/marcosfrancomarinho/go-first-project/src/shared/exceptions"
)

type GormCreatorUser struct{}

func NewGormCreatorUser() interfaces.CreateUser {
	return &GormCreatorUser{}
}

func (g *GormCreatorUser) Create(user *entities.UserRegister) error {
	client := database.NewDataBase()

	datas := database.User{
		Id:       user.GetID(),
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}

	if result := client.Create(&datas); result.Error != nil {
		return exceptions.ErrUserAlreadyExists
	}
	return nil
}
