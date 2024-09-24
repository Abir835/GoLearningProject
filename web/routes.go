package web

import (
	"go-learning-project/web/handlers"
	"go-learning-project/web/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /health-check",
		manager.With(
			http.HandlerFunc(handlers.HealthCheck),
		),
	)
}
