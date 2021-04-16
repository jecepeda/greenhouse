package greenhouse

import "github.com/jecepeda/greenhouse/server/gutil"

func FakeMonitoringData(deviceID uint64) MonitoringData {
	return MonitoringData{
		DeviceID:          deviceID,
		Temperature:       gutil.RandFloat(0, 50),
		Humidity:          gutil.RandFloat(0, 100),
		HeaterEnabled:     gutil.RandBool(),
		HumidifierEnabled: gutil.RandBool(),
	}
}
