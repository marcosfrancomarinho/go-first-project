package entities

import "github.com/marcosfrancomarinho/src/domain/valuesobject"

type UserRegister struct {
	name     *valuesobject.Name
	email    *valuesobject.Email
	password *valuesobject.Password
	id       *valuesobject.ID
}

func NewUserRegister(name *valuesobject.Name, email *valuesobject.Email, password *valuesobject.Password, id *valuesobject.ID) (*UserRegister, error) {
	return &UserRegister{name: name, email: email, password: password, id: id}, nil
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
