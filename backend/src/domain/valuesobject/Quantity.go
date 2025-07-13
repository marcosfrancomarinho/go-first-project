package valuesobject

import "errors"

type Quantity struct {
	quantity int
}

func NewQuantity(quantity int) (*Quantity, error) {
	if err := validateQuantity(quantity); err != nil {
		return nil, err
	}
	return &Quantity{quantity: quantity}, nil
}

func validateQuantity(quantity int) error {
	if quantity <= 0 {
		return errors.New("quantidade nao pode ser igual a zeror ou menor")
	}
	return nil
}
func (q *Quantity) GetValue() int {
	return q.quantity
}
