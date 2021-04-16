package greenhousehandler_test

import (
	"net/http"
	"testing"

	"github.com/jecepeda/greenhouse/server/domain/greenhouse/greenhousehandler"
	"github.com/jecepeda/greenhouse/server/gtest"
	"github.com/jecepeda/greenhouse/server/gutil"
	"github.com/stretchr/testify/suite"
)

type GreenhouseHandlerTestSuite struct {
	gtest.GTestSuite
}

func (suite *GreenhouseHandlerTestSuite) TestSaveMonitoringData() {
	at := gtest.SetupMockedAtomic(suite.Pool, suite.Tx)
	defer gtest.VerifySucceeded(suite.T(), at)

	d := suite.CreateDevice()

	request := greenhousehandler.SaveMonitoringDataRequest{
		Temperature:       gutil.RandFloat(0, 50),
		Humidity:          gutil.RandFloat(0, 100),
		HeaterEnabled:     gutil.RandBool(),
		HumidifierEnabled: gutil.RandBool(),
	}

	tt := gtest.TestRequest{
		Body:     gtest.MarshalJSON(request),
		DeviceID: d.ID,
		Method:   "POST",
	}.Run(greenhousehandler.SaveMonitoringData(suite.DC))

	suite.Require().Equal(http.StatusOK, tt.Code)
}

func TestGreenhouseHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(GreenhouseHandlerTestSuite))
}
