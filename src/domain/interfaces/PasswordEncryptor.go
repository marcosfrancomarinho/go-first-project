package interfaces

import v "github.com/marcosfrancomarinho/src/domain/valuesobject"

type PasswordEncryptor interface {
	Encryptor(password *v.Password) (*v.Password, error)
}
