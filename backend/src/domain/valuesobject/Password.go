package valuesobject

import (
	"errors"
	"regexp"
	"strings"
)

type Password struct {
	password string
}

func NewPassword(password string) (*Password, error) {
	trimmed := strings.TrimSpace(password)

	if err := validatePassword(password); err != nil {
		return nil, err
	}
	return &Password{password: trimmed}, nil
}

func validatePassword(password string) error {
	if len(password) == 0 {
		return errors.New("senha obrigatoria")
	}
	regexMinLen := regexp.MustCompile(`^.{8,}$`)
	regexLetterUpperCase := regexp.MustCompile(`[A-Z]`)
	regexHasNumber := regexp.MustCompile(`\d`)
	regexHasSymbol := regexp.MustCompile(`[^a-zA-Z0-9]`)

	if !regexMinLen.MatchString(password) {
		return errors.New("senha tem que ter no minimo oito digito")
	}
	if !regexLetterUpperCase.MatchString(password) {
		return errors.New("senha tem pelo menos uma letra maiscula")
	}
	if !regexHasNumber.MatchString(password) {
		return errors.New("senha tem que ter pelo menos um numero")
	}
	if !regexHasSymbol.MatchString(password) {
		return errors.New("senha tem que ter pelo menos um simbolo")
	}
	return nil
}

func (p *Password) GetValue() string {
	return p.password
}

func (p *Password) With(password string) *Password {
	return &Password{password: password}
}
