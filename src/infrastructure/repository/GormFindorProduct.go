package repository

import (
	"errors"

	"github.com/marcosfrancomarinho/go-first-project/src/application/mappers"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
)

type GormFindorProduct struct{}

func (f *GormFindorProduct) FindAll() (*[]entities.Product, error) {
	var (
		client               database.Database
		products             []database.Product
		FindorProductMappers mappers.FindorProductMappers
	)
	client.Connection()

	result := client.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("nenhum  produto encontrado")
	}
	listProducts, err := FindorProductMappers.GetAllProducts(&products)
	if err != nil {
		return nil, err
	}
	return listProducts, nil
}
