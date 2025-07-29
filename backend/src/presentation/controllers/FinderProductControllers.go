package controllers

import (
	"strconv"
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
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
	const PAGE string = "page"
	const LIMIT string = "limit"

	queries, err := httpContext.GetQuery(PAGE, LIMIT)
	if err != nil {
		httpContext.SendError(err)
		return
	}
	page, err := strconv.Atoi(queries[PAGE])
	if err != nil {
		httpContext.SendError(err)
		return
	}
	limit, err := strconv.Atoi(queries[LIMIT])
	if err != nil {
		httpContext.SendError(err)
		return
	}
	input := dto.RequestFindorProductDTO{Page: page, Limit: limit}
	output, err := f.findorProductUseCase.FindAll(&input)
	if err != nil {
		httpContext.SendError(err)
		return
	}
	httpContext.Send(200, output)
}
