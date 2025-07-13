package valuesobject

import (
	"errors"
	"strings"
)

type ID struct {
	id string
}

func NewID(id string) (*ID, error) {
	trimmed := strings.TrimSpace(id)
	if err := validateID(id); err != nil {
		return nil, err
	}
	return &ID{id: trimmed}, nil
}

func validateID(id string) error {
	if len(id) == 0 {
		return errors.New("id é obrigatorio")
	}
	return nil
}

func (i *ID) GetValue() string {
	return i.id
}
