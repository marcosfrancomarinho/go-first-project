package entities

import "github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"

type Product struct {
	id       *valuesobject.ID
	name     *valuesobject.Name
	price    *valuesobject.Price
	quantity *valuesobject.Quantity
}

func NewProduct(
	id *valuesobject.ID,
	name *valuesobject.Name,
	price *valuesobject.Price,
	quantity *valuesobject.Quantity,
) *Product {
	return &Product{id, name, price, quantity}
}

func (p *Product) GetName() string {
	return p.name.GetValue()
}
func (p *Product) GetID() string {
	return p.id.GetValue()
}

func (p *Product) GetPrice() float64 {
	return p.price.GetValue()
}
func (p *Product) GetQuantity() int {
	return p.quantity.GetValue()
}

func (p *Product) GetTotal() float64 {
	return p.GetPrice() * float64(p.GetQuantity())
}
