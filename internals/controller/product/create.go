package product

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/product"
	api_v1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
	"log"
)

func (c *Controller) Create(ctx context.Context, request *api_v1.CreateRequest) (*empty.Empty, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.product.Create",
	)
	defer span.Finish()
	log.Println("SPAN", span)
	span.LogKV("Input Handler", request)
	_, err := c.service.Create(&model.Request{
		Name:   request.GetName(),
		Brand:  request.GetBrand(),
		Detail: request.GetDetail(),
		Price:  request.GetPrice(),
	})

	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}
