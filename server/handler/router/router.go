package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router represents a set of routes from
// the same domain, and all the paths that correspond to that domain
type Router struct {
	Handler http.HandlerFunc
	Paths   []string
}

// Endpoint represents the necessary data to build
// router
type Endpoint struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

// NewRouterFromEndpoints builds a router from a set of endpoints
func NewRouterFromEndpoints(endpoints []Endpoint) Router {
	r := mux.NewRouter()
	paths := make([]string, len(endpoints))
	for i, e := range endpoints {
		r.HandleFunc(e.Path, e.HandlerFunc).Methods(e.Method)
		paths[i] = e.Path
	}
	return Router{
		Handler: r.ServeHTTP,
		Paths:   paths,
	}
}

// AddToMux adds the router to mux router
func AddToMux(m *mux.Router, r Router) {
	for _, p := range r.Paths {
		m.HandleFunc(p, r.Handler)
	}
}
