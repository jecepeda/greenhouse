// Package gsql contains the necessary tools to run
// transactions inside this repo, based on interfaces
package gsql

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// Common represents the basic interface
// for making queries. It can be run either
// inside or outside a transaction.
type Common interface {
	sqlx.Ext
	sqlx.ExtContext
	sqlx.Queryer
	sqlx.QueryerContext
	sqlx.Preparer
	sqlx.PreparerContext
}

// Atomic gathers and transaction
// and makes common operations inside
// it, such as finishing or rollbacking a transaction
type Atomic interface {
	GetTx() *sqlx.Tx
	End() error
	Fail() error
}

// TransactionPool manages transactions
// inside the codebase, so when a service
// needs to init a transaction, calls this service
type TransactionPool interface {
	Start(ctx context.Context) (Atomic, error)
}
