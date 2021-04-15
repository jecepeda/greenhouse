package device

import (
	"context"

	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindByID(ctx context.Context, deviceID uint64, db gsql.Common) (Device, error)
	FindAll(ctx context.Context, db gsql.Common) ([]Device, error)
	StoreDevice(ctx context.Context, tx *sqlx.Tx, name string, pass []byte) (Device, error)
}

type Service interface {
	FindByID(ctx context.Context, deviceID uint64) (Device, error)
	SaveDevice(ctx context.Context, name, password string) (Device, error)
}
