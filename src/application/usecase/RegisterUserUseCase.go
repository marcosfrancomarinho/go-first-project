package usecase

import (
	D "github.com/marcosfrancomarinho/src/application/DTO"
	E "github.com/marcosfrancomarinho/src/domain/entities"
	I "github.com/marcosfrancomarinho/src/domain/interfaces"
	V "github.com/marcosfrancomarinho/src/domain/valuesobject"
)

type RegisterUserUseCase struct {
	createUser I.CreateUser
	gerator    I.GeratorID
	encryptor  I.PasswordEncryptor
}

func NewRegisterUserUseCase(createUser I.CreateUser, gerator I.GeratorID, encryptor I.PasswordEncryptor) *RegisterUserUseCase {
	return &RegisterUserUseCase{createUser: createUser, gerator: gerator, encryptor: encryptor}
}

func (r *RegisterUserUseCase) Register(userDTO *D.RequestRegisterUserDTO) (*D.ResponseRegisterUserDTO, error) {

	name, err := V.NewName(userDTO.Name)
	if err != nil {
		return nil, err
	}

	email, err := V.NewEmail(userDTO.Email)
	if err != nil {
		return nil, err
	}

	password, err := V.NewPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}

	id, err := r.gerator.Generate()
	if err != nil {
		return nil, err
	}

	userRegister, err := E.NewUserRegister(name, email, password, id)
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

	return &D.ResponseRegisterUserDTO{
		Message: "usuario criado com sucesso.",
		IdUser:  id.GetValue(),
	}, nil

}
