package entities

import V "github.com/marcosfrancomarinho/src/domain/valuesobject"

type UserRegister struct {
	name     *V.Name
	email    *V.Email
	password *V.Password
	id       *V.ID
}

func NewUserRegister(name *V.Name, email *V.Email, password *V.Password, id *V.ID) (*UserRegister, error) {
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

func (u *UserRegister) UpdatePassword(password *V.Password) {
	u.password = password
}
