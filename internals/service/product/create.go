package product

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/product"
)

func (s *ProcessTracingService) Create(request *model.Request) (int, error) {
	input := &entity.Product{
		Name:   request.Name,
		Detail: request.Detail,
		Brand:  request.Brand,
		Price:  request.Price,
	}

	err := s.repository.Create(input)

	return input.ID, err
}
