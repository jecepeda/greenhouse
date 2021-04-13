package gsql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Transaction struct {
	transaction *sqlx.Tx
}

func (t *Transaction) GetTx() *sqlx.Tx {
	return t.transaction
}

func (t *Transaction) End() error {
	if t.transaction == nil {
		return nil
	}
	if err := t.transaction.Commit(); err != nil {
		return errors.Wrap(err, "could not end transaction")
	}
	return nil
}

func (t *Transaction) Fail() error {
	if t.transaction == nil {
		return nil
	}
	if err := t.transaction.Rollback(); err != nil {
		return errors.Wrap(err, "could not end transaction")
	}
	return nil
}
