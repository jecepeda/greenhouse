package gsql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Pool struct {
	db *sqlx.DB
}

func (p *Pool) Start(ctx context.Context) (Atomic, error) {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not start transaction")
	}
	return &Transaction{transaction: tx}, nil
}
