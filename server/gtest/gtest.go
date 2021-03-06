// Package gtest contains the utilities
// to test inside the server package
package gtest

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/icrowley/fake"
	"github.com/jecepeda/greenhouse/server/crypt"
	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/dcontainer"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// GTestSuite is the main structure used in our tests
// to gather all setup/teardown/utility logic into one place
// so we don't need to write the same code everywhere
type GTestSuite struct {
	suite.Suite
	DC   handler.DependencyContainer
	DB   *sqlx.DB
	Tx   *sqlx.Tx
	Pool *gsql.MockTransactionPool
}

func (suite *GTestSuite) SetupSuite() {
	var err error
	suite.DB, err = sqlx.Open("postgres", os.Getenv("SQL_DB"))
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}
}

func (suite *GTestSuite) SetupTest() {
	var err error
	ctx := context.Background()
	suite.Tx, err = suite.DB.BeginTxx(ctx, nil)
	suite.Require().NoError(err)
	suite.Pool = &gsql.MockTransactionPool{}
	suite.DC = dcontainer.NewDependencyContainer(suite.Tx)
	suite.DC.SetEncrypter(crypt.BEncrypter{})
	suite.DC.SetTransactionPool(suite.Pool)
	suite.DC.Init()
}

func (suite *GTestSuite) TearDownTest() {
	suite.Require().NoError(suite.Tx.Rollback())
}

func (suite *GTestSuite) CreateDevice() device.Device {
	name := fake.CharactersN(20)
	password := fake.CharactersN(20)
	ctx := context.Background()

	d, err := suite.DC.GetDeviceService().SaveDevice(ctx, name, password)
	suite.Require().NoError(err)
	return d
}

// SetupMockedAtomic setups the transaction pool so when
// a service function inits a transaction, we return a mock atomic
func SetupMockedAtomic(pool *gsql.MockTransactionPool, tx *sqlx.Tx) *gsql.AtomicMock {
	at := gsql.NewAtomicMock(tx)
	at.On("End").Return()
	at.On("Fail").Return()
	pool.On("Start", mock.Anything).Return(at, nil)
	return at
}

// VerifySucceeded checks that the atomic function End has been called
// and the Fail didn't
func VerifySucceeded(t *testing.T, at *gsql.AtomicMock) {
	t.Helper()
	at.AssertCalled(t, "End")
	at.AssertNotCalled(t, "Fail")
}
