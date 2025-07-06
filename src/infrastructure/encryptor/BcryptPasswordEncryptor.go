package encryptor

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordEncryptor struct {
}

func (b *BcryptPasswordEncryptor) Encryptor(password *valuesobject.Password) (*valuesobject.Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password.GetValue()), bcrypt.DefaultCost)
	hashPassword := password.With(string(bytes))
	return hashPassword, err

}

func (b *BcryptPasswordEncryptor) ValidatePassword(password string, encryptedPassword string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(encryptedPassword),
		[]byte(password),
	)
}
