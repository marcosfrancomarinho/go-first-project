package encryptor

import (
	"github.com/marcosfrancomarinho/src/domain/valuesobject"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordEncryptor struct {
}

func (b *BcryptPasswordEncryptor) Encryptor(password *valuesobject.Password) (*valuesobject.Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password.GetValue()), bcrypt.DefaultCost)
	hashPassword := password.With(string(bytes))
	return hashPassword, err

}
