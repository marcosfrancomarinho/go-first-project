package controllers

import (
	"github.com/marcosfrancomarinho/src/application/dto"
	"github.com/marcosfrancomarinho/src/application/usecase"
	"github.com/marcosfrancomarinho/src/domain/interfaces"
)

type RegisterUserControllers struct {
	registerUserUseCase *usecase.RegisterUserUseCase
}

func NewRegisterUserControllers(registerUserUseCase *usecase.RegisterUserUseCase) *RegisterUserControllers {
	return &RegisterUserControllers{registerUserUseCase: registerUserUseCase}
}

type Raw struct {
	Name     string
	Email    string
	Password string
}

func (r *RegisterUserControllers) Execute(httpContext interfaces.HttpContext) {
	var raw Raw
	err := httpContext.GetBody(&raw)
	if err != nil {
		httpContext.SendError(err)
		return
	}

	input := dto.RequestRegisterUserDTO{Name: raw.Name, Email: raw.Email, Password: raw.Password}

	output, err := r.registerUserUseCase.Register(&input)
	if err != nil {
		httpContext.SendError(err)
		return
	}
	httpContext.Send(201, output)
}
