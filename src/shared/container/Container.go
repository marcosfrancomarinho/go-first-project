package container

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/auth"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/encryptor"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/geradorid"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/repository"
	"github.com/marcosfrancomarinho/go-first-project/src/presentation/controllers"
)

type Container struct{}

var instance *Container

func GetInstance() *Container {
	if instance == nil {
		instance = &Container{}
	}
	return instance
}

type Controllers struct {
	RegisterUser interfaces.HttpControllers
	LoginUser    interfaces.HttpControllers
}

func (c *Container) Dependencies() *Controllers {
	encryptor := encryptor.BcryptPasswordEncryptor{}

	creatorUser := repository.GormCreatorUser{}
	idGerador := geradorid.UUID{}
	registerUserUseCase := usecase.NewRegisterUserUseCase(&creatorUser, &idGerador, &encryptor)
	registerUserControllers := controllers.NewRegisterUserControllers(registerUserUseCase)

	findorUser := repository.GormFindorUser{}
	userAuthenticator := auth.JwtUserAuthenticator{}
	loginUserUseCase := usecase.NewLoginUserUseCase(&encryptor, &userAuthenticator, &findorUser)
	loginUserControllers := controllers.NewLoginUserControllers(loginUserUseCase)

	return &Controllers{
		RegisterUser: registerUserControllers,
		LoginUser:    loginUserControllers,
	}
}
