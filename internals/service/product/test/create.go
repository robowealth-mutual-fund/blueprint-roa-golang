package test

func (suite *PackageTestSuite) TestCreate() {
	givenInput := suite.makeTestCreateInput()
	givenTracing := suite.makeTestTracing()
	suite.repo.On("Create", givenTracing).Once().Return(nil)
	_, err := suite.service.Create(givenInput)
	suite.NoError(err)
}
