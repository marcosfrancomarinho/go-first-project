package valuesobject

import (
	"errors"
	"regexp"
	"strings"
)

type Email struct {
	email string
}

func NewEmail(email string) (*Email, error) {
	trimmed := strings.TrimSpace(email)
	if err := validateEmail(trimmed); err != nil {
		return nil, err
	}
	return &Email{email: trimmed}, nil
}

func (e *Email) GetValue() string {
	return e.email
}

func validateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("email é obrigatorio")
	}

	isValid := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if !isValid.MatchString(email) {
		return errors.New("e-mail inválido. Use um formato como exemplo@dominio.com")
	}
	
	return nil
}
