package controllers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type CreatorProductControllers struct {
	creatorProductUseCase *usecase.CreatorProductUseCase
}

func NewCreatorProductControllers(creatorProductUseCase *usecase.CreatorProductUseCase) interfaces.HttpControllers {
	return &CreatorProductControllers{creatorProductUseCase}
}

type RawCreatorProduct struct {
	Name     string
	Price    float32
	Quantity int
}

func (c *CreatorProductControllers) Execute(httpContext interfaces.HttpContext) {
	var raw RawCreatorProduct
	if err := httpContext.GetBody(&raw); err != nil {
		httpContext.SendError(err)
		return
	}

	input := dto.RequestCreatorProductDTO{Name: raw.Name, Price: raw.Price, Quantity: raw.Quantity}

	output, err := c.creatorProductUseCase.Create(&input)
	if err != nil {
		httpContext.SendError(err)
		return
	}

	httpContext.Send(200, output)
}
