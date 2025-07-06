package controllers

import (
	D "github.com/marcosfrancomarinho/src/application/DTO"
	U "github.com/marcosfrancomarinho/src/application/usecase"
	I "github.com/marcosfrancomarinho/src/domain/interfaces"
)

type RegisterUserControllers struct {
	registerUserUseCase *U.RegisterUserUseCase
}

func NewRegisterUserControllers(registerUserUseCase *U.RegisterUserUseCase) *RegisterUserControllers {
	return &RegisterUserControllers{registerUserUseCase: registerUserUseCase}
}

type Raw struct {
	Name     string
	Email    string
	Password string
}

func (r *RegisterUserControllers) Execute(httpContext I.HttpContext) {
	var raw Raw
	err := httpContext.GetBody(&raw)
	if err != nil {
		httpContext.SendError(err)
		return
	}

	input := D.RequestRegisterUserDTO{Name: raw.Name, Email: raw.Email, Password: raw.Password}

	output, err := r.registerUserUseCase.Register(&input)
	if err != nil {
		httpContext.SendError(err)
		return
	}

	httpContext.Send(201, output.ToObject())
}
