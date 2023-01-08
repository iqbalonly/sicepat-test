package api

import (
	"net/http"
	"sicepat/internal/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// BuildRouter builds the HTTP router for this service
func BuildRouter(storage storage.StorageDAO) *mux.Router {
	router := mux.NewRouter()

	addNotFoundHandler(router)

	healthCheckAPI := &healthCheck{}
	healthCheckAPI.AddRoutes(router)

	userAPI := NewUserHandler(storage)
	userAPI.AddRoutes(router)

	addMiddlewares(router)

	return router
}

func addMiddlewares(router *mux.Router) {
	// defines middleware here
}

// instruments and handles unknown request paths
func addNotFoundHandler(router *mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		logrus.Warn("route not found", logrus.Fields{
			"path":   req.URL.EscapedPath(),
			"method": req.Method,
		})

		resp.WriteHeader(http.StatusNotFound)
		_, _ = resp.Write([]byte("Where are u going ?"))
	})
}
