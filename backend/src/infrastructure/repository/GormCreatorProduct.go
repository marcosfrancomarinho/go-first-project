package repository

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
)

type GormCreatorProduct struct {
}

func NewGormCreatorProduct() interfaces.CreatorProduct {
	return &GormCreatorProduct{}
}




func (g *GormCreatorProduct) Create(product *entities.Product) error {
	var client database.Database

	client.Connection()

	datas := database.Product{
		Id:       product.GetID(),
		Name:     product.GetName(),
		Price:    product.GetPrice(),
		Quantity: product.GetQuantity(),
	}
	if result := client.DB.Create(&datas); result.Error != nil {
		return result.Error
	}
	return nil
}
