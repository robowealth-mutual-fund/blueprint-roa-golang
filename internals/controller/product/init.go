package product

import service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/product"

type Controller struct {
	service service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{
		service: s,
	}
}
