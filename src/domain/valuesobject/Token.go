package valuesobject

import (
	"errors"
	"strings"
)

type Token struct {
	token string
}

func NewToken(token string) (*Token, error) {
	trimmed := strings.TrimSpace(token)
	if err := validateToken(token); err != nil {
		return nil, err
	}
	return &Token{token: trimmed}, nil
}

func validateToken(token string) error {
	if len(token) == 0 {
		return errors.New("token é obrigatorio")
	}
	return nil
}

func (t *Token) GetValue() string {
	return t.token
}
