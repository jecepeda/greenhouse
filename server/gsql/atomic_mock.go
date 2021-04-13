package gsql

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type AtomicMock struct {
	mock.Mock
	tx *sqlx.Tx
}

func NewAtomicMock(tx *sqlx.Tx) *AtomicMock {
	return &AtomicMock{
		tx: tx,
	}
}

// End is called when COMMIT transaction
func (a *AtomicMock) End() error {
	a.Called()
	return nil
}

// Fail is called when ROLLBACK transaction
func (a *AtomicMock) Fail() error {
	a.Called()
	return nil
}

// GetTx returns the current transaction
func (a *AtomicMock) GetTx() *sqlx.Tx {
	return a.tx
}
