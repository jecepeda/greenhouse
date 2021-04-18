package greenhouse

import "errors"

var (
	// ErrMonitoringDataNotFound is thrown when the monitoring data is not found on the
	// database
	ErrMonitoringDataNotFound = errors.New("monitoring data not found")
)
