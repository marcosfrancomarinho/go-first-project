package container

import (
	U "github.com/marcosfrancomarinho/src/application/usecase"
	I "github.com/marcosfrancomarinho/src/domain/interfaces"
	E "github.com/marcosfrancomarinho/src/infrastructure/encryptor"
	G "github.com/marcosfrancomarinho/src/infrastructure/geradorid"
	R "github.com/marcosfrancomarinho/src/infrastructure/repository"
	C "github.com/marcosfrancomarinho/src/presentation/controllers"
)

type Container struct{}

func (c *Container) Dependecies() map[string]I.HttpControllers {

	creatorUser := R.GormCreatorUser{}
	idGerador := G.UUID{}
	encryptor := E.BcryptPasswordEncryptor{}
	creatorUserUseCase := U.NewRegisterUserUseCase(&creatorUser, &idGerador, &encryptor)
	creatorUserControllers := C.NewRegisterUserControllers(creatorUserUseCase)

	return map[string]I.HttpControllers{
		"creator-user": creatorUserControllers,
	}
}
