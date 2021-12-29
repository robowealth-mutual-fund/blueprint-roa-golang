package product_test

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/product/test"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
