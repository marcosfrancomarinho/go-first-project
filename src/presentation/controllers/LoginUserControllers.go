package controllers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type LoginUserControllers struct {
	loginUserUseCase *usecase.LoginUserUseCase
}

type RawLogin struct {
	Email    string
	Password string
}

func NewLoginUserControllers(loginUserUseCase *usecase.LoginUserUseCase) *LoginUserControllers {
	return &LoginUserControllers{loginUserUseCase: loginUserUseCase}
}

func (l *LoginUserControllers) Execute(httpContext interfaces.HttpContext) {
	var raw RawLogin
	if err := httpContext.GetBody(&raw); err != nil {
		httpContext.SendError(err)
		return
	}
	input := dto.RequestLoginUserDTO{Email: raw.Email, Password: raw.Password}

	output, err := l.loginUserUseCase.Login(&input)
	if err != nil {
		httpContext.SendError(err)
		return
	}

	httpContext.Send(200, output, output.Token)

}
