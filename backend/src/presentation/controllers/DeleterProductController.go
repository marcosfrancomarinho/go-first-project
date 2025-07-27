package controllers

import (
	"fmt"

	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type RawDeleterProduct struct {
	Id string
}

type DeleterProductController struct {
	deleterProductUseCase *usecase.DeleterProductUseCase
}

func NewDeleterProductController(deleterProductUseCase *usecase.DeleterProductUseCase) interfaces.HttpControllers {
	return &DeleterProductController{deleterProductUseCase}
}

func (d *DeleterProductController) Execute(httpContext interfaces.HttpContext) {
	var raw RawDeleterProduct

	id, err := httpContext.GetParams("id")

	if err != nil {
		httpContext.SendError(err)
		return
	}
	fmt.Println(raw)
	input := &dto.RequestDeleterProductDTO{Id: *id}

	output, err := d.deleterProductUseCase.Delete(input)
	if err != nil {
		httpContext.SendError(err)
		return
	}
	httpContext.Send(200, output)
}
