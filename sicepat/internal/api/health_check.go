package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// healthCheck is a standard, simple health check
type healthCheck struct{}

// AddRoutes adds the routers for this API to the provided router (or subrouter)
func (h *healthCheck) AddRoutes(router *mux.Router) {
	router.HandleFunc("/health", h.handler).Methods(http.MethodGet)
	router.HandleFunc("/", h.handler).Methods(http.MethodGet)
}

func (h *healthCheck) handler(resp http.ResponseWriter, _ *http.Request) {
	_, _ = resp.Write([]byte(`OK`))
}
