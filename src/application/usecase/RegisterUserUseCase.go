package usecase

import (
	"github.com/marcosfrancomarinho/src/application/dto"
	"github.com/marcosfrancomarinho/src/domain/entities"
	"github.com/marcosfrancomarinho/src/domain/interfaces"
	"github.com/marcosfrancomarinho/src/domain/valuesobject"
)

type RegisterUserUseCase struct {
	createUser interfaces.CreateUser
	gerator    interfaces.GeratorID
	encryptor  interfaces.PasswordEncryptor
}

func NewRegisterUserUseCase(createUser interfaces.CreateUser, gerator interfaces.GeratorID, encryptor interfaces.PasswordEncryptor) *RegisterUserUseCase {
	return &RegisterUserUseCase{createUser: createUser, gerator: gerator, encryptor: encryptor}
}

func (r *RegisterUserUseCase) Register(userDTO *dto.RequestRegisterUserDTO) (*dto.ResponseRegisterUserDTO, error) {

	name, err := valuesobject.NewName(userDTO.Name)
	if err != nil {
		return nil, err
	}

	email, err := valuesobject.NewEmail(userDTO.Email)
	if err != nil {
		return nil, err
	}

	password, err := valuesobject.NewPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}

	id, err := r.gerator.Generate()
	if err != nil {
		return nil, err
	}

	userRegister, err := entities.NewUserRegister(name, email, password, id)
	if err != nil {
		return nil, err
	}

	encryptedPassword, err := r.encryptor.Encryptor(password)
	if err != nil {
		return nil, err
	}

	userRegister.UpdatePassword(encryptedPassword)

	err = r.createUser.Create(userRegister)
	if err != nil {
		return nil, err
	}

	return &dto.ResponseRegisterUserDTO{
		Message: "usuario criado com sucesso.",
		IdUser:  id.GetValue(),
	}, nil

}
