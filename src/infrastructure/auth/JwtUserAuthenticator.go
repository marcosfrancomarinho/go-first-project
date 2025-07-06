package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
	"os"
	"time"
)

type JwtUserAuthenticator struct{}

type claims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func getKeyEnv(key string) ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	keySecret := os.Getenv(key)
	return []byte(keySecret), nil
}

func (j *JwtUserAuthenticator) GenerateToken(user *entities.UserRegister) (*valuesobject.Token, error) {
	expirationTime := time.Now().Add(60 * time.Minute)

	jwtKey, err := getKeyEnv("KEY_SECRET")
	if err != nil {
		return nil, err
	}

	claimsOptions := &claims{
		UserId: user.GetID(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "app-user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsOptions)

	hash, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}
	createdToken, err := valuesobject.NewToken(hash)
	if err != nil {
		return nil, err
	}
	return createdToken, nil
}

func (j *JwtUserAuthenticator) ValidateToken(token *valuesobject.Token) error {
	return errors.New("method no implement")
}
