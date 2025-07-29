package usecase

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type FindorProductUseCase struct {
	FindorProduct interfaces.FindorProduct
}

func NewFindorProductUseCase(FindorProduct interfaces.FindorProduct) *FindorProductUseCase {
	return &FindorProductUseCase{FindorProduct: FindorProduct}
}

func (f *FindorProductUseCase) FindAll(payload *dto.RequestFindorProductDTO) (*dto.ResponseFindorProductDTO, error) {
	var listProducts []dto.Products

	limit, err := valuesobject.NewLimit(payload.Limit)
	if err != nil {
		return nil, err
	}

	page, err := valuesobject.NewPage(payload.Page)
	if err != nil {
		return nil, err
	}
	containerPage := entities.NewContainerPage(limit, page)

	products, err := f.FindorProduct.FindAll(containerPage)
	if err != nil {
		return nil, err
	}
	
	total, err := f.FindorProduct.TotalCount()
	if err != nil {
		return nil, err
	}

	for _, value := range *products {
		listProducts = append(listProducts,
			dto.Products{
				Id:       value.GetID(),
				Name:     value.GetName(),
				Price:    value.GetPrice(),
				Quantity: value.GetQuantity(),
				Total:    value.GetTotal(),
			})
	}

	return &dto.ResponseFindorProductDTO{Total: total.GetValue(), Products: listProducts}, nil
}
