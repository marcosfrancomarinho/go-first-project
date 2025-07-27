package usecase

import (
	"fmt"

	"github.com/marcosfrancomarinho/go-first-project/src/application/dto"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type DeleterProductUseCase struct {
	deleterProduct interfaces.DeleterProduct
}

func NewDeleterProductUseCase(deleterProduct interfaces.DeleterProduct) *DeleterProductUseCase {
	return &DeleterProductUseCase{deleterProduct}
}

func (d *DeleterProductUseCase) Delete(payload *dto.RequestDeleterProductDTO) (*dto.ResponseDeleterProductDTO, error) {
	id, err := valuesobject.NewID(payload.Id)
	if err != nil {
		return nil, err
	}
	if err := d.deleterProduct.Delete(id); err != nil {
		return nil, err
	}
	return &dto.ResponseDeleterProductDTO{
		Id:      id.GetValue(),
		Message: fmt.Sprintf("product Id: %s deletado com sucesso", id.GetValue()),
	}, nil
}
