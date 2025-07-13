package controllers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type FinderProductControllers struct {
	finderProductUseCase *usecase.FinderProductUseCase
}

func NewFinderProductControllers(finderProductUseCase *usecase.FinderProductUseCase) *FinderProductControllers {
	return &FinderProductControllers{finderProductUseCase: finderProductUseCase}
}

func (f *FinderProductControllers) Execute(httpContext interfaces.HttpContext) {
	output, err := f.finderProductUseCase.FindAll()
	if err != nil {
		httpContext.SendError(err)
		return
	}
	httpContext.Send(200, output)
}
