package greenhouse

import (
	"context"

	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jmoiron/sqlx"
)

// Repository contains the persistence functions and all the logic to save/extract the monitoring data
// into/from the database
type Repository interface {
	// StoreMonitoringData stores monitoring data into the database
	StoreMonitoringData(ctx context.Context, tx *sqlx.Tx, data MonitoringData) (MonitoringData, error)
	// FindByID gets monitoring data given an id
	FindByID(ctx context.Context, db gsql.Common, id uint64) (MonitoringData, error)
}

// Service contains the service functions and the logic to save/extract monitoring data into/from the database
type Service interface {
	// SaveMonitoringData stores monitoring data into the database, dealing with transactions
	SaveMonitoringData(ctx context.Context, data MonitoringData) (MonitoringData, error)
	// FindByID finds the monitoring data given the id
	FindByID(ctx context.Context, id uint64) (MonitoringData, error)
}
