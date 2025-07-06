package encryptor

import (
	V "github.com/marcosfrancomarinho/src/domain/valuesobject"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordEncryptor struct {
}

func (b *BcryptPasswordEncryptor) Encryptor(password *V.Password) (*V.Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password.GetValue()), bcrypt.DefaultCost)
	hashPassword := password.With(string(bytes))
	return hashPassword, err

}
