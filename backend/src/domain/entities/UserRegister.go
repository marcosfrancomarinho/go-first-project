package entities

import "github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"

type UserRegister struct {
	name     *valuesobject.Name
	email    *valuesobject.Email
	password *valuesobject.Password
	id       *valuesobject.ID
}

func NewUserRegister(
	name *valuesobject.Name,
	email *valuesobject.Email,
	password *valuesobject.Password,
	id *valuesobject.ID,
) *UserRegister {
	return &UserRegister{name, email, password, id}
}

func (u *UserRegister) GetName() string {
	return u.name.GetValue()
}
func (u *UserRegister) GetEmail() string {
	return u.email.GetValue()
}
func (u *UserRegister) GetPassword() string {
	return u.password.GetValue()
}
func (u *UserRegister) GetID() string {
	return u.id.GetValue()
}

func (u *UserRegister) UpdatePassword(password *valuesobject.Password) {
	u.password = password
}
