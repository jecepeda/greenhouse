package devicerepo

import (
	"context"
	"database/sql"
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

func (r *Repository) FindByID(ctx context.Context, deviceID uint64, db gsql.Common) (device.Device, error) {
	var (
		errMsg = "find by id"
		query  = `SELECT * from devices where id = $1`
		d      device.Device
	)

	err := db.QueryRowxContext(ctx, query, deviceID).StructScan(&d)
	if errors.Is(err, sql.ErrNoRows) {
		return device.Device{}, errors.Wrap(device.ErrNotFound, errMsg)
	} else if err != nil {
		return device.Device{}, errors.Wrap(err, errMsg)
	}

	return d, nil
}

func (r *Repository) StoreDevice(ctx context.Context, tx *sqlx.Tx, name string, password []byte) (device.Device, error) {
	var (
		errMsg = "store device"
		query  = `INSERT INTO devices(name, password, created_at, updated_at)
				  VALUES (:name, :password, :created_at, :updated_at) RETURNING (id)`
	)
	created := device.Device{
		Name:      name,
		Password:  password,
		CreatedAt: time.Now().Round(time.Microsecond).UTC(),
		UpdatedAt: time.Now().Round(time.Microsecond).UTC(),
	}

	var deviceID uint64
	rows, err := sqlx.NamedQueryContext(ctx, tx, query, created)
	if err != nil {
		return device.Device{}, errors.Wrap(err, errMsg)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&deviceID)
		if err != nil {
			return device.Device{}, errors.Wrap(err, errMsg)
		}
		created.ID = deviceID
	}

	return created, nil
}
