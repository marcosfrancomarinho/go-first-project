package valuesobject

import (
	"errors"
	"strings"
)

type Name struct {
	name string
}

func NewName(name string) (*Name, error) {
	trimmed := strings.TrimSpace(name)
	if err := validateName(name); err != nil {
		return nil, err
	}
	return &Name{name: trimmed}, nil
}

func (n *Name) GetValue() string {
	return n.name
}

func validateName(name string) error {
	if len(name) == 0 {
		return errors.New("nome obrigatorio")
	}

	// isValid := regexp.MustCompile(`^[A-Za-zÀ-ÖØ-öø-ÿ]+(?: [A-Za-zÀ-ÖØ-öø-ÿ]+)*$`)

	// if !isValid.MatchString(name) {
	// 	return errors.New("o nome deve conter apenas letras e espaços sem números ou símbolos")
	// }

	return nil
}
