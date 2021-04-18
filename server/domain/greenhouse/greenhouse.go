// Package greenhouse contains all telemetry logic,
// including persistence, service functions and handlers.
package greenhouse

import "time"

// MonitoringData represents the data that a device can
// monitor right now
type MonitoringData struct {
	ID                uint64    `json:"id" db:"id"`
	DeviceID          uint64    `json:"device_id" db:"device_id"`
	Temperature       float64   `json:"temperature" db:"temperature"`
	Humidity          float64   `json:"humidity" db:"humidity"`
	HeaterEnabled     bool      `json:"heater_enabled" db:"heater_enabled"`
	HumidifierEnabled bool      `json:"humidifier_enabled" db:"humidifier_enabled"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
}

// AsTest clones the monitoring data with some modified fields
// so they can be checked on comparison functions without giving errors
func (m MonitoringData) AsTest() MonitoringData {
	newData := m
	newData.CreatedAt = newData.CreatedAt.Round(time.Microsecond).UTC()
	return newData
}
