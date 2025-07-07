package repository

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
)

type GormCreatorUser struct{}

func (g *GormCreatorUser) Create(user *entities.UserRegister) error {
	var client database.Database

	client.Connection()

	datas := database.User{
		Id:       user.GetID(),
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}

	if result := client.DB.Create(&datas); result.Error != nil {
		return result.Error
	}
	return nil
}
