package test

import (
	"context"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	processTracingModel "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/product"
	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/product"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils/mocks"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx     context.Context
	repo    *mocks.Repository
	service *service.ProcessTracingService
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repo = &mocks.Repository{}
	suite.service = service.NewService(suite.repo)
}

func (suite *PackageTestSuite) makeTestTracing() (verbose *entity.Product) {
	//now := carbon.Now().Timestamp()
	return &entity.Product{
		PlatformID:  "odini",
		AccountID:   "odini1234",
		ProcessName: "openacc",
		Level:       "in",
		State:       "register",
	}
}

func (suite *PackageTestSuite) makeTestCreateInput() (input *processTracingModel.Request) {
	return &processTracingModel.Request{
		PlatformID:  "odini",
		AccountID:   "odini1234",
		ProcessName: "openacc",
		Level:       "in",
		State:       "register",
	}
}
