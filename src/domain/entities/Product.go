package entities

import "github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"

type Product struct {
	Id       *valuesobject.ID
	Name     *valuesobject.Name
	Price    *valuesobject.Price
	Quantity *valuesobject.Quantity
}

func NewProduct(
	Id *valuesobject.ID,
	Name *valuesobject.Name,
	Price *valuesobject.Price,
	Quantity *valuesobject.Quantity,
) *Product {
	return &Product{Id: Id, Name: Name, Price: Price, Quantity: Quantity}
}

func (p *Product) GetName() string {
	return p.Name.GetValue()
}
func (p *Product) GetID() string {
	return p.Id.GetValue()
}

func (p *Product) GetPrice() float32 {
	return p.Price.GetValue()
}
func (p *Product) GetQuantity() int {
	return p.Quantity.GetValue()
}
