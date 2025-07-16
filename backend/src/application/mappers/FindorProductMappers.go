package mappers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
)

type FindorProductMappers struct {
}

func NewFindorProductMappers() *FindorProductMappers {
	return &FindorProductMappers{}
}

func (f *FindorProductMappers) GetAllProducts(products *[]database.Product) (*[]entities.Product, error) {
	var listProducts []entities.Product

	for _, value := range *products {
		id, err := valuesobject.NewID(value.Id)
		if err != nil {
			return nil, err
		}
		name, err := valuesobject.NewName(value.Name)
		if err != nil {
			return nil, err
		}
		price, err := valuesobject.NewPrice(value.Price)
		if err != nil {
			return nil, err
		}
		quantity, err := valuesobject.NewQuantity(value.Quantity)
		if err != nil {
			return nil, err
		}

		listProducts = append(listProducts, *entities.NewProduct(id, name, price, quantity))
	}
	return &listProducts, nil
}
