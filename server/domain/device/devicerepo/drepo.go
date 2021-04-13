package devicerepo

import (
	"context"
	"time"

	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct{}

func (r *Repository) FindAll(ctx context.Context, db gsql.Common) ([]device.Device, error) {
	var (
		errMsg  = "find all"
		query   = `SELECT * from devices`
		devices []device.Device
	)

	err := sqlx.SelectContext(ctx, db, &devices, query)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}

	return devices, nil
}

func (r *Repository) StoreDevice(ctx context.Context, tx *sqlx.Tx, name string, password []byte) (device.Device, error) {
	var (
		errMsg = "store device"
		query  = `INSERT INTO devices(name, password, created_at, updated_at)
				  VALUES (:name, :password, :created_at, :updated_at)`
	)
	created := device.Device{
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, err := tx.NamedExecContext(ctx, query, created)
	if err != nil {
		return device.Device{}, errors.Wrap(err, errMsg)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return device.Device{}, errors.Wrap(err, errMsg)
	}
	created.ID = uint(id)
	return created, nil
}
