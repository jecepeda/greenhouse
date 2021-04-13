package device

import (
	"context"

	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindAll(ctx context.Context, db gsql.Common) ([]Device, error)
	StoreDevice(ctx context.Context, tx *sqlx.Tx, name string, pass []byte) (Device, error)
}

type Service interface {
	SaveDevice(ctx context.Context, name, password string) (Device, error)
}
