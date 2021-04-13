package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Handler http.HandlerFunc
	Paths   []string
}

type Spec struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

func NewRouterFromEndpoints(endpoints []Spec) Router {
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

func AddToMux(m *mux.Router, r Router) {
	for _, p := range r.Paths {
		m.HandleFunc(p, r.Handler)
	}
}
