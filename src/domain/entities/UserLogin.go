package entities

import "github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"

type UserLogin struct {
	email    valuesobject.Email
	password valuesobject.Password
}

func NewUserLogin(email valuesobject.Email, password valuesobject.Password, id valuesobject.ID) (*UserLogin, error) {
	return &UserLogin{email: email, password: password}, nil
}
