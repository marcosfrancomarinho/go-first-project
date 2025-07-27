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
	"sync"
)

type Container struct{}

var (
	instance *Container
	once     sync.Once
)

func GetInstance() *Container {
	once.Do(func() {
		instance = &Container{}
	})
	return instance
}

type Handlers struct {
	RegisterUserControllers      interfaces.HttpControllers
	LoginUserControllers         interfaces.HttpControllers
	UserAuthenticatorMiddlewares interfaces.HttpControllers
	CreatorProductControllers    interfaces.HttpControllers
	FindorProductControllers     interfaces.HttpControllers
	DeleterProductController     interfaces.HttpControllers
}

func (c *Container) Dependencies() *Handlers {
	encryptor := encryptor.NewBcryptPasswordEncryptor()
	userAuthenticator := auth.NewJwtUserAuthenticator()
	idGerador := geradorid.NewUUID()

	creatorUser := repository.NewGormCreatorUser()
	registerUserUseCase := usecase.NewRegisterUserUseCase(creatorUser, idGerador, encryptor)
	registerUserControllers := controllers.NewRegisterUserControllers(registerUserUseCase)

	findorUser := repository.NewGormFindorUser()
	loginUserUseCase := usecase.NewLoginUserUseCase(encryptor, userAuthenticator, findorUser)
	loginUserControllers := controllers.NewLoginUserControllers(loginUserUseCase)

	userAuthenticatorMiddlewares := middlewares.NewUserAuthenticatorMiddlewares(userAuthenticator)

	creatorProduct := repository.NewGormCreatorProduct()
	creatorProductUsecase := usecase.NewCreatorProductUseCase(creatorProduct, idGerador)
	creatorProductControllers := controllers.NewCreatorProductControllers(creatorProductUsecase)

	findorProduct := repository.NewGormFindorProduct()
	findorProductUseCase := usecase.NewFindorProductUseCase(findorProduct)
	findorProductControllers := controllers.NewFindorProductControllers(findorProductUseCase)

	deleterProduct := repository.NewGormDeleterProduct()
	deleterProductUseCase := usecase.NewDeleterProductUseCase(deleterProduct)
	deleterProductController := controllers.NewDeleterProductController(deleterProductUseCase)

	return &Handlers{
		registerUserControllers,
		loginUserControllers,
		userAuthenticatorMiddlewares,
		creatorProductControllers,
		findorProductControllers,
		deleterProductController,
	}
}
