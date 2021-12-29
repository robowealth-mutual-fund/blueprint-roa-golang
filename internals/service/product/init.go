package product

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils"
)

type ProcessTracingService struct {
	repository utils.Repository
}

func NewService(r utils.Repository) (service *ProcessTracingService) {
	return &ProcessTracingService{
		repository: r,
	}
}
