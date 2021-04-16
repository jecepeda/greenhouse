package greenhouse

import (
	"context"

	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	StoreMonitoringData(ctx context.Context, tx *sqlx.Tx, data MonitoringData) (MonitoringData, error)
	FindByID(ctx context.Context, db gsql.Common, id uint64) (MonitoringData, error)
}

type Service interface {
	SaveMonitoringData(ctx context.Context, data MonitoringData) (MonitoringData, error)
	FindByID(ctx context.Context, id uint64) (MonitoringData, error)
}
