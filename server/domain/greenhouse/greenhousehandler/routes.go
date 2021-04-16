package greenhousehandler

import (
	"github.com/jecepeda/greenhouse/server/auth"
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/router"
)

func NewAuthRouter(dc handler.DependencyContainer) router.Router {
	endpoints := []router.Endpoint{
		{
			Path:        "/v1/monitoring/data",
			Method:      "POST",
			HandlerFunc: auth.MatchDeviceID(SaveMonitoringData(dc)),
		},
	}
	return router.NewRouterFromEndpoints(endpoints)
}
