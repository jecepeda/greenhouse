package devicehandler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/icrowley/fake"
	"github.com/jecepeda/greenhouse/server/domain/device/devicehandler"
	"github.com/jecepeda/greenhouse/server/gtest"
	"github.com/stretchr/testify/suite"
)

type DeviceHandlerTestSuite struct {
	gtest.GTestSuite
}

func (suite *DeviceHandlerTestSuite) TestAuthentication() {
	name := fake.CharactersN(20)
	password := fake.CharactersN(20)
	ctx := context.Background()

	at := gtest.SetupMockedAtomic(suite.Pool, suite.Tx)
	defer gtest.VerifySucceeded(suite.T(), at)

	d, err := suite.DC.GetDeviceService().SaveDevice(ctx, name, password)
	suite.Require().NoError(err)

	suite.Run("test login", func() {
		data := map[string]interface{}{
			"device":   d.ID,
			"password": password,
		}

		tt := gtest.TestRequest{
			Body:   gtest.MarshalJSON(data),
			Method: "POST",
			IsForm: true,
		}.Run(devicehandler.Login(suite.DC))

		suite.Require().Equal(http.StatusOK, tt.Code)

		response := make(map[string]string)
		suite.Require().NoError(gtest.UnMarshalJSON(tt, &response))
		_, ok := response["access_token"]
		suite.True(ok)
		_, ok = response["refresh_token"]
		suite.True(ok)
	})

	suite.Run("test refresh", func() {
		tt := gtest.TestRequest{
			Body:      nil,
			Method:    "POST",
			DeviceID:  d.ID,
			IsRefresh: true,
		}.Run(devicehandler.Refresh(suite.DC))

		suite.Require().Equal(http.StatusOK, tt.Code)
	})
}

func TestDeviceHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceHandlerTestSuite))
}
