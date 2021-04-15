package gtest

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jecepeda/greenhouse/server/crypt"
	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/dcontainer"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

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

func SetupMockedAtomic(pool *gsql.MockTransactionPool, tx *sqlx.Tx) *gsql.AtomicMock {
	at := gsql.NewAtomicMock(tx)
	at.On("End").Return()
	at.On("Fail").Return()
	pool.On("Start", mock.Anything).Return(at, nil)
	return at
}

func VerifySucceeded(t *testing.T, at *gsql.AtomicMock) {
	t.Helper()
	at.AssertCalled(t, "End")
	at.AssertNotCalled(t, "Fail")
}
