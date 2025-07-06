package repository

import (
	E "github.com/marcosfrancomarinho/src/domain/entities"
	D "github.com/marcosfrancomarinho/src/infrastructure/database"
)

type GormCreatorUser struct{}

func (g *GormCreatorUser) Create(user *E.UserRegister) error {
	var db D.Database
	db.Connection()
	datas := D.User{
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
