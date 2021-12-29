//go:build integration
// +build integration

package test

func (suite *PackageTestSuite) TestCreate() {
	input := suite.makeTestStruct("odini", "odini01121", "opem", "in", "Test")
	err := suite.repo.Create(input)
	suite.NoError(err)
}
