package interfaces

import "github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"

type PasswordEncryptor interface {
	Encryptor(password *valuesobject.Password) (*valuesobject.Password, error)
	ValidatePassword(password string, encryptedPassword string) error
}
