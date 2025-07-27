package repository

import (
	"fmt"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/database"
)

type GormDeleterProduct struct {
}

func NewGormDeleterProduct() interfaces.DeleterProduct {
	return &GormDeleterProduct{}
}
func (d *GormDeleterProduct) Delete(id *valuesobject.ID) error {
	client := database.NewDataBase()
	product := database.Product{Id: id.GetValue()}

	result := client.Delete(&product)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("product id %s nao cadastrado", id.GetValue())
	}
	return nil
}
