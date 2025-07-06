package container

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/usecase"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/encryptor"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/geradorid"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/repository"
	"github.com/marcosfrancomarinho/go-first-project/src/presentation/controllers"
)

type Container struct{}

type Controllers struct {
	CreatorUser interfaces.HttpControllers
}

func (c *Container) Dependencies() *Controllers {

	creatorUser := repository.GormCreatorUser{}
	idGerador := geradorid.UUID{}
	encryptor := encryptor.BcryptPasswordEncryptor{}
	creatorUserUseCase := usecase.NewRegisterUserUseCase(&creatorUser, &idGerador, &encryptor)
	creatorUserControllers := controllers.NewRegisterUserControllers(creatorUserUseCase)

	return &Controllers{
		CreatorUser: creatorUserControllers,
	}
}
