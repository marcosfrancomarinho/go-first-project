package valuesobject

import "errors"

type Price struct {
	price float64
}

func NewPrice(price float64) (*Price, error) {
	if err := validatePrice(price); err != nil {
		return nil, err
	}
	return &Price{price: price}, nil
}

func validatePrice(price float64) error {
	if  price <= 0 {
		return errors.New("preço não pode ser 0 ou menor do que zero")
	}
	return nil
}

func (p *Price) GetValue() float64 {
	return p.price
}
