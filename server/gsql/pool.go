package gsql

import (
	"context"

	"github.com/jecepeda/greenhouse/server/gerror"
	"github.com/jmoiron/sqlx"
)

type Pool struct {
	DB *sqlx.DB
}

func NewPool(db *sqlx.DB) Pool {
	return Pool{
		DB: db,
	}
}

func (p Pool) Start(ctx context.Context) (Atomic, error) {
	tx, err := p.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, gerror.Wrap(err, "could not start transaction")
	}
	return &Transaction{transaction: tx}, nil
}
