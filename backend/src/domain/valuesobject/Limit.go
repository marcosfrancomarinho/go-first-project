package valuesobject

import "errors"

type Limit struct {
	limit int
}

func NewLimit(limit int) (*Limit, error) {
	if err := validateLimit(limit); err != nil {
		return nil, err
	}
	return &Limit{limit}, nil
}

func validateLimit(limit int) error {
	if limit <= 0 {
		return errors.New("lmite de produto não pode ser 0 ou menor do que zero")
	}
	return nil
}

func (l *Limit) GetValue() int {
	return l.limit
}
