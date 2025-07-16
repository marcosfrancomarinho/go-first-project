package repository

import (
	"errors"

	"github.com/marcosfrancomarinho/go-first-project/src/application/mappers"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
)

type GormFindorProduct struct{}

func NewGormFindorProduct() interfaces.FindorProduct {
	return &GormFindorProduct{}
}

func (f *GormFindorProduct) FindAll() (*[]entities.Product, error) {
	client := database.NewDataBase()
	products := &[]database.Product{}

	findorProductMappers := mappers.NewFindorProductMappers()

	result := client.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("nenhum  produto encontrado")
	}
	listProducts, err := findorProductMappers.GetAllProducts(products)
	if err != nil {
		return nil, err
	}
	return listProducts, nil
}
