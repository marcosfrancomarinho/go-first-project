package entities

import "github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"

type UserLogin struct {
	email    *valuesobject.Email
	password *valuesobject.Password
}

func NewUserLogin(email *valuesobject.Email, password *valuesobject.Password) *UserLogin {
	return &UserLogin{email, password}
}

func (u *UserLogin) GetEmail() string {
	return u.email.GetValue()
}
func (u *UserLogin) GetPassword() string {
	return u.password.GetValue()
}
