package greenhousehandler

import (
	"context"
	"net/http"

	"github.com/jecepeda/greenhouse/server/domain/greenhouse"
	"github.com/jecepeda/greenhouse/server/gerror"
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/httputil"
)

type SaveMonitoringDataRequest struct {
	Temperature       float64 `json:"temperature"`
	Humidity          float64 `json:"humidity"`
	HeaterEnabled     bool    `json:"heater_enabled"`
	HumidifierEnabled bool    `json:"humidifier_enabled"`
}

// SaveMonitoringData saves monitoring data
func SaveMonitoringData(dc handler.DependencyContainer) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var (
			errMsg   = "save monitoring data"
			deviceID uint64
			data     SaveMonitoringDataRequest
		)
		ctx, cancel := context.WithTimeout(r.Context(), handler.DefaultDuration)
		defer cancel()
		if err := httputil.GetIDFromURL(r, "deviceID", &deviceID); err != nil {
			http.Error(rw, gerror.Wrap(err, errMsg).Error(), http.StatusBadRequest)
			return
		}
		if err := httputil.ReadJSON(r, &data); err != nil {
			http.Error(rw, gerror.Wrap(err, errMsg).Error(), http.StatusBadRequest)
			return
		}
		monitoringData := greenhouse.MonitoringData{
			DeviceID:          deviceID,
			Temperature:       data.Temperature,
			Humidity:          data.Humidity,
			HeaterEnabled:     data.HeaterEnabled,
			HumidifierEnabled: data.HumidifierEnabled,
		}
		savedData, err := dc.GetGreenhouseService().SaveMonitoringData(ctx, monitoringData)
		if err != nil {
			http.Error(rw, gerror.Wrap(err, errMsg).Error(), http.StatusInternalServerError)
			return
		}
		httputil.WriteJSON(rw, savedData)
	}
}
