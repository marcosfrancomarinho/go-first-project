package usecase

import (
	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type LoginUserUseCase struct {
	encryptor         interfaces.PasswordEncryptor
	userAuthenticator interfaces.UserAuthenticator
	findorUser        interfaces.FindorUser
}

func NewLoginUserUseCase(
	encryptor interfaces.PasswordEncryptor,
	userAuthenticator interfaces.UserAuthenticator,
	findorUser interfaces.FindorUser,
) *LoginUserUseCase {
	return &LoginUserUseCase{encryptor: encryptor, userAuthenticator: userAuthenticator, findorUser: findorUser}
}

func (l *LoginUserUseCase) Login(userDTO *dto.RequestLoginUserDTO) (*dto.ResponseLoginUserDTO, error) {
	email, err := valuesobject.NewEmail(userDTO.Email)
	if err != nil {
		return nil, err
	}
	password, err := valuesobject.NewPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}
	userLogin, err := entities.NewUserLogin(email, password)
	if err != nil {
		return nil, err
	}
	foundUser, err := l.findorUser.FindByEmail(userLogin)
	if err != nil {
		return nil, err
	}
	if err := l.encryptor.ValidatePassword(password.GetValue(), foundUser.GetPassword()); err != nil {
		return nil, err
	}
	token, err := l.userAuthenticator.GenerateToken(foundUser)
	if err != nil {
		return nil, err
	}
	return &dto.ResponseLoginUserDTO{
		IdUser:  foundUser.GetID(),
		Message: "usuario logado com sucesso",
		Token:   token.GetValue(),
	}, nil
}
