package usecase

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type RegisterUserUseCase struct {
	createUser interfaces.CreateUser
	gerator    interfaces.GeratorID
	encryptor  interfaces.PasswordEncryptor
}

func NewRegisterUserUseCase(
	createUser interfaces.CreateUser,
	gerator interfaces.GeratorID,
	encryptor interfaces.PasswordEncryptor,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{createUser: createUser, gerator: gerator, encryptor: encryptor}
}

func (r *RegisterUserUseCase) Register(payload *dto.RequestRegisterUserDTO) (*dto.ResponseRegisterUserDTO, error) {

	name, err := valuesobject.NewName(payload.Name)
	if err != nil {
		return nil, err
	}

	email, err := valuesobject.NewEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	password, err := valuesobject.NewPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	id, err := r.gerator.Generate()
	if err != nil {
		return nil, err
	}

	userRegister := entities.NewUserRegister(name, email, password, id)

	encryptedPassword, err := r.encryptor.Encryptor(password)
	if err != nil {
		return nil, err
	}

	userRegister.UpdatePassword(encryptedPassword)

	if err = r.createUser.Create(userRegister); err != nil {
		return nil, err
	}

	return &dto.ResponseRegisterUserDTO{
		Message: "usuario criado com sucesso.",
		IdUser:  id.GetValue(),
	}, nil

}
