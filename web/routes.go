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

	mux.Handle(
		"GET /books",
		manager.With(
			http.HandlerFunc(handlers.GetBooks),
		),
	)

	mux.Handle(
		"GET /book",
		manager.With(
			http.HandlerFunc(handlers.GetBookById),
		),
	)

	mux.Handle(
		"POST /book",
		manager.With(
			http.HandlerFunc(handlers.InsertBook),
		),
	)

	mux.Handle(
		"PUT /book",
		manager.With(
			http.HandlerFunc(handlers.UpdateBookById),
		),
	)

	mux.Handle(
		"DELETE /book",
		manager.With(
			http.HandlerFunc(handlers.DeleteBookById),
		),
	)
}
