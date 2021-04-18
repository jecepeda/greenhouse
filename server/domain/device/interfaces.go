package device

import (
	"context"

	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jmoiron/sqlx"
)

// Repository manages the persistence layer
type Repository interface {
	// FindByID finds a device given its id
	FindByID(ctx context.Context, deviceID uint64, db gsql.Common) (Device, error)
	// FindAll finds all devices on the database
	FindAll(ctx context.Context, db gsql.Common) ([]Device, error)
	// StoreDevice stores a device inside the database
	StoreDevice(ctx context.Context, tx *sqlx.Tx, name string, pass []byte) (Device, error)
}

type Service interface {
	// FindByID finds a device given its id
	FindByID(ctx context.Context, deviceID uint64) (Device, error)
	// SaveDevice saves a device on the persistence layer
	SaveDevice(ctx context.Context, name, password string) (Device, error)
}
