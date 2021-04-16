package devicehandler

import (
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/router"
)

func NewRouter(dc handler.DependencyContainer) router.Router {
	endpoints := []router.Endpoint{
		{
			Path:        "/v1/device/login",
			Method:      "POST",
			HandlerFunc: Login(dc),
		},
		{
			Path:        "/v1/device/refresh",
			Method:      "POST",
			HandlerFunc: Refresh(dc),
		},
	}
	return router.NewRouterFromEndpoints(endpoints)
}
