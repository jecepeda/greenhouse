package greenhouserepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jecepeda/greenhouse/server/domain/greenhouse"
	"github.com/jecepeda/greenhouse/server/gerror"
	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jmoiron/sqlx"
)

type Repository struct{}

// StoreMonitoringData stores monitoring data into the database
func (r *Repository) StoreMonitoringData(ctx context.Context, tx *sqlx.Tx, data greenhouse.MonitoringData) (greenhouse.MonitoringData, error) {
	var (
		errMsg = "store monitoring data"
		query  = `INSERT INTO
					monitoring_data(device_id, temperature, humidity, heater_enabled, humidifier_enabled, created_at)
					VALUES(:device_id, :temperature, :humidity, :heater_enabled, :humidifier_enabled, :created_at)
					RETURNING (id)`
	)
	data.CreatedAt = time.Now().Round(time.Microsecond).UTC()

	var id uint64
	rows, err := sqlx.NamedQueryContext(ctx, tx, query, data)
	if err != nil {
		return greenhouse.MonitoringData{}, gerror.Wrap(err, errMsg)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return greenhouse.MonitoringData{}, gerror.Wrap(err, errMsg)
		}
		data.ID = id
	}

	return data, nil
}

func (r *Repository) FindByID(ctx context.Context, db gsql.Common, id uint64) (greenhouse.MonitoringData, error) {
	var (
		errMsg = "find by id"
		query  = "SELECT * FROM monitoring_data where id = $1"
		data   greenhouse.MonitoringData
	)

	err := db.QueryRowxContext(ctx, query, id).StructScan(&data)
	if errors.Is(err, sql.ErrNoRows) {
		return greenhouse.MonitoringData{}, gerror.Wrap(greenhouse.ErrMonitoringDataNotFound, errMsg)
	} else if err != nil {
		return greenhouse.MonitoringData{}, gerror.Wrap(err, errMsg)
	}

	return data, nil
}
