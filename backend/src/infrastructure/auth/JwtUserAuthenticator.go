package auth

import (
	"errors"
	"fmt"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/entities"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type JwtUserAuthenticator struct{}

func NewJwtUserAuthenticator() interfaces.UserAuthenticator {
	return &JwtUserAuthenticator{}
}

type claims struct {
	IdUser string `json:"userId"`
	jwt.RegisteredClaims
}

func getKeyEnv(key string) ([]byte, error) {
	if os.Getenv("GIN_MODE") != "release" {
		_ = godotenv.Load()
	}

	keySecret := os.Getenv(key)

	if keySecret == "" {
		return nil, fmt.Errorf("variável de ambiente %s não definida", key)
	}

	return []byte(keySecret), nil
}

func (j *JwtUserAuthenticator) GenerateToken(user *entities.UserRegister) (*valuesobject.Token, error) {
	expirationTime := time.Now().Add(60 * time.Minute)

	jwtKey, err := getKeyEnv("KEY_SECRET")
	if err != nil {
		return nil, err
	}

	claimsOptions := &claims{
		IdUser: user.GetID(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
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

func (j *JwtUserAuthenticator) ValidateToken(token *valuesobject.Token) (*valuesobject.ID, error) {
	jwtKey, err := getKeyEnv("KEY_SECRET")
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.ParseWithClaims(token.GetValue(), &claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("token invalido")
	}

	claims, ok := parsedToken.Claims.(*claims)
	if !ok {
		return nil, errors.New("falha ao converter claims")
	}

	idUser, err := valuesobject.NewID(claims.IdUser)
	if err != nil {
		return nil, err
	}

	return idUser, nil
}
