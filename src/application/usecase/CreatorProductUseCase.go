package usecase

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type CreatorProductUseCase struct {
	creatorProduct interfaces.CreatorProduct
	geratorId      interfaces.GeratorID
}

func NewCreatorProductUseCase(
	creatorProduct interfaces.CreatorProduct,
	geratorId interfaces.GeratorID,
) *CreatorProductUseCase {
	return &CreatorProductUseCase{creatorProduct: creatorProduct, geratorId: geratorId}
}

func (c *CreatorProductUseCase) Create(payload *dto.RequestCreatorProductDTO) (*dto.ResponseCreatorProductDTO, error) {
	name, err := valuesobject.NewName(payload.Name)
	if err != nil {
		return nil, err
	}

	price, err := valuesobject.NewPrice(payload.Price)
	if err != nil {
		return nil, err
	}

	quantity, err := valuesobject.NewQuantity(payload.Quantity)
	if err != nil {
		return nil, err
	}

	id, err := c.geratorId.Generate()
	if err != nil {
		return nil, err
	}

	product := entities.NewProduct(id, name, price, quantity)

	if err := c.creatorProduct.Create(product); err != nil {
		return nil, err
	}
	
	return &dto.ResponseCreatorProductDTO{
		Id:      id.GetValue(),
		Message: "produto criado com sucesso",
	}, nil

}
