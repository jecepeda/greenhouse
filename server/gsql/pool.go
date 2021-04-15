package gsql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
		return nil, errors.Wrap(err, "could not start transaction")
	}
	return &Transaction{transaction: tx}, nil
}
