package gsql

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Common interface {
	sqlx.Ext
	sqlx.ExtContext
	sqlx.Queryer
	sqlx.QueryerContext
	sqlx.Preparer
	sqlx.PreparerContext
}

type Atomic interface {
	GetTx() *sqlx.Tx
	End() error
	Fail() error
}

type TransactionPool interface {
	Start(ctx context.Context) (Atomic, error)
}
