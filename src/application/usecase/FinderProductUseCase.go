package usecase

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type FinderProductUseCase struct {
	finderProduct interfaces.FinderProduct
}

func NewFinderProductUseCase(finderProduct interfaces.FinderProduct) *FinderProductUseCase {
	return &FinderProductUseCase{finderProduct: finderProduct}
}

func (f *FinderProductUseCase) FindAll() (*[]dto.ResponseFinderProductDTO, error) {
	var listProducts []dto.ResponseFinderProductDTO
	products, err := f.finderProduct.FindAll()
	if err != nil {
		return nil, err
	}
	for _, value := range *products {
		listProducts = append(listProducts,
			dto.ResponseFinderProductDTO{
				Id:       value.GetID(),
				Name:     value.GetName(),
				Price:    value.GetPrice(),
				Quantity: value.GetQuantity(),
				Total:    value.GetTotal(),
			})
	}
	return &listProducts, nil
}
