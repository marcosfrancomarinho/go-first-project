package repository

import (
	"github.com/marcosfrancomarinho/src/domain/entities"
	"github.com/marcosfrancomarinho/src/infrastructure/database"
)

type GormCreatorUser struct{}

func (g *GormCreatorUser) Create(user *entities.UserRegister) error {
	var db database.Database
	db.Connection()
	datas := database.User{
		Id:       user.GetID(),
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}
	result := db.DB.Create(&datas)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
