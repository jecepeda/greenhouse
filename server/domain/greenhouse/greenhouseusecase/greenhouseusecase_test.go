package greenhouseusecase_test

import (
	"context"
	"testing"

	"github.com/jecepeda/greenhouse/server/domain/greenhouse"
	"github.com/jecepeda/greenhouse/server/gtest"
	"github.com/stretchr/testify/suite"
)

type GreenhouseUsecaseTest struct {
	gtest.GTestSuite
}

func (suite *GreenhouseUsecaseTest) TestStoreMonitoringData() {
	at := gtest.SetupMockedAtomic(suite.Pool, suite.Tx)
	defer gtest.VerifySucceeded(suite.T(), at)

	d := suite.CreateDevice()

	data := greenhouse.FakeMonitoringData(d.ID)
	data, err := suite.DC.GetGreenhouseService().SaveMonitoringData(context.Background(), data)
	suite.Require().NoError(err)

	got, err := suite.DC.GetGreenhouseService().FindByID(context.Background(), data.ID)
	suite.Require().NoError(err)

	suite.Equal(data.AsTest(), got.AsTest())
}

func TestGreenhouseUsecaseTest(t *testing.T) {
	suite.Run(t, new(GreenhouseUsecaseTest))
}
