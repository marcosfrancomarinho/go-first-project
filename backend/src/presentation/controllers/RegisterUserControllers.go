package controllers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type RegisterUserControllers struct {
	registerUserUseCase *usecase.RegisterUserUseCase
}

func NewRegisterUserControllers(registerUserUseCase *usecase.RegisterUserUseCase) interfaces.HttpControllers {
	return &RegisterUserControllers{registerUserUseCase}
}

type RawRegister struct {
	Name     string
	Email    string
	Password string
}

func (r *RegisterUserControllers) Execute(httpContext interfaces.HttpContext) {
	var raw RawRegister
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
