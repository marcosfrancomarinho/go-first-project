package controllers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)
type RawFinderProduct struct{
	Page int `form:"page"`
	Limit int `form:"limit"`
}

type FindorProductControllers struct {
	findorProductUseCase *usecase.FindorProductUseCase
}

func NewFindorProductControllers(findorProductUseCase *usecase.FindorProductUseCase) interfaces.HttpControllers {
	return &FindorProductControllers{findorProductUseCase}
}

func (f *FindorProductControllers) Execute(httpContext interfaces.HttpContext) {
	var raw RawFinderProduct
	if err := httpContext.GetQuery(&raw); err != nil {
		httpContext.SendError(err)
		return
	}	
	input := dto.RequestFindorProductDTO{Page: raw.Page, Limit: raw.Limit}
	output, err := f.findorProductUseCase.FindAll(&input)
	if err != nil {
		httpContext.SendError(err)
		return
	}
	httpContext.Send(200, output)
}
