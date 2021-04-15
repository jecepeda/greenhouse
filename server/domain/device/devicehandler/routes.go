package devicehandler

import (
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/router"
)

func GetRoutes(dc handler.DependencyContainer) router.Router {
	endpoints := []router.Spec{
		{
			Path:        "/v1/device/login",
			Method:      "GET",
			HandlerFunc: login(dc),
		},
	}
	return router.NewRouterFromEndpoints(endpoints)
}
