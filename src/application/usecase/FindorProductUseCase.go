package usecase

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type FindorProductUseCase struct {
	FindorProduct interfaces.FindorProduct
}

func NewFindorProductUseCase(FindorProduct interfaces.FindorProduct) *FindorProductUseCase {
	return &FindorProductUseCase{FindorProduct: FindorProduct}
}

func (f *FindorProductUseCase) FindAll() (*[]dto.ResponseFindorProductDTO, error) {
	var listProducts []dto.ResponseFindorProductDTO
	products, err := f.FindorProduct.FindAll()
	if err != nil {
		return nil, err
	}
	for _, value := range *products {
		listProducts = append(listProducts,
			dto.ResponseFindorProductDTO{
				Id:       value.GetID(),
				Name:     value.GetName(),
				Price:    value.GetPrice(),
				Quantity: value.GetQuantity(),
				Total:    value.GetTotal(),
			})
	}
	return &listProducts, nil
}
