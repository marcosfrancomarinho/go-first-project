package container

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/auth"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/encryptor"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/geradorid"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/repository"
	"github.com/marcosfrancomarinho/go-first-project/src/presentation/controllers"
	"github.com/marcosfrancomarinho/go-first-project/src/presentation/middlewares"
)

type Container struct{}

var instance *Container

func GetInstance() *Container {
	if instance == nil {
		instance = &Container{}
	}
	return instance
}

type handlers struct {
	RegisterUserControllers      interfaces.HttpControllers
	LoginUserControllers         interfaces.HttpControllers
	UserAuthenticatorMiddlewares interfaces.HttpControllers
	CreatorProductControllers    interfaces.HttpControllers
	FinderProductControllers     interfaces.HttpControllers
}

func (c *Container) Dependencies() *handlers {
	encryptor := encryptor.BcryptPasswordEncryptor{}
	userAuthenticator := auth.JwtUserAuthenticator{}
	idGerador := geradorid.UUID{}

	creatorUser := repository.GormCreatorUser{}
	registerUserUseCase := usecase.NewRegisterUserUseCase(&creatorUser, &idGerador, &encryptor)
	registerUserControllers := controllers.NewRegisterUserControllers(registerUserUseCase)

	findorUser := repository.GormFindorUser{}
	loginUserUseCase := usecase.NewLoginUserUseCase(&encryptor, &userAuthenticator, &findorUser)
	loginUserControllers := controllers.NewLoginUserControllers(loginUserUseCase)

	userAuthenticatorMiddlewares := middlewares.NewUserAuthenticatorMiddlewares(&userAuthenticator)

	creatorProduct := repository.GormCreatorProduct{}
	creatorProductUsecase := usecase.NewCreatorProductUseCase(&creatorProduct, &idGerador)
	creatorProductControllers := controllers.NewCreatorProductControllers(creatorProductUsecase)

	finderProduct := repository.GormFinderProduct{}
	finderProductUseCase := usecase.NewFinderProductUseCase(&finderProduct)
	finderProductControllers := controllers.NewFinderProductControllers(finderProductUseCase)

	return &handlers{
		RegisterUserControllers:      registerUserControllers,
		LoginUserControllers:         loginUserControllers,
		UserAuthenticatorMiddlewares: userAuthenticatorMiddlewares,
		CreatorProductControllers:    creatorProductControllers,
		FinderProductControllers:     finderProductControllers,
	}
}
