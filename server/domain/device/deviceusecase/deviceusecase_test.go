package deviceusecase_test

import (
	"context"
	"testing"

	_ "github.com/lib/pq"

	"github.com/icrowley/fake"
	"github.com/jecepeda/greenhouse/server/gtest"
	"github.com/stretchr/testify/suite"
)

type DeviceUseCaseTest struct {
	gtest.GTestSuite
}

func (suite *DeviceUseCaseTest) TestSaveDevice() {
	name := fake.CharactersN(20)
	password := fake.CharactersN(20)
	ctx := context.Background()

	at := gtest.SetupMockedAtomic(suite.Pool, suite.Tx)
	defer gtest.VerifySucceeded(suite.T(), at)

	d, err := suite.DC.GetDeviceService().SaveDevice(ctx, name, password)
	suite.Require().NoError(err)
	suite.NotZero(d.ID)
	suite.Len(d.Password, 60)

	got, err := suite.DC.GetDeviceService().FindByID(ctx, d.ID)
	suite.Require().NoError(err)
	suite.Equal(got.AsTest(), d.AsTest())
}

func TestDeviceUseCase(t *testing.T) {
	suite.Run(t, new(DeviceUseCaseTest))
}
