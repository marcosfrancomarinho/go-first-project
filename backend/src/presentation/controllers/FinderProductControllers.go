package controllers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type FindorProductControllers struct {
	findorProductUseCase *usecase.FindorProductUseCase
}

func NewFindorProductControllers(findorProductUseCase *usecase.FindorProductUseCase) interfaces.HttpControllers {
	return &FindorProductControllers{findorProductUseCase}
}

func (f *FindorProductControllers) Execute(httpContext interfaces.HttpContext) {
	output, err := f.findorProductUseCase.FindAll()
	if err != nil {
		httpContext.SendError(err)
		return
	}
	httpContext.Send(200, output)
}
