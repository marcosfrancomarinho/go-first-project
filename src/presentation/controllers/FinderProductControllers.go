package controllers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type FindorProductControllers struct {
	FindorProductUseCase *usecase.FindorProductUseCase
}

func NewFindorProductControllers(FindorProductUseCase *usecase.FindorProductUseCase) *FindorProductControllers {
	return &FindorProductControllers{FindorProductUseCase: FindorProductUseCase}
}

func (f *FindorProductControllers) Execute(httpContext interfaces.HttpContext) {
	output, err := f.FindorProductUseCase.FindAll()
	if err != nil {
		httpContext.SendError(err)
		return
	}
	httpContext.Send(200, output)
}
