package entities

import V "github.com/marcosfrancomarinho/src/domain/valuesobject"

type UserLogin struct {
	email    V.Email
	password V.Password
}

func NewUserLogin(email V.Email, password V.Password, id V.ID) (*UserLogin, error) {
	return &UserLogin{email: email, password: password}, nil
}
