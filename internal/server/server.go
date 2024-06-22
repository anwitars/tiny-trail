package server

import (
	"net/http"
	"tinytrail/internal/endpoints"
	"tinytrail/internal/middlewares"

	"github.com/jmoiron/sqlx"
)

type AppContext struct {
	DB *sqlx.DB
}

func registerEndpoint(pattern string, handler http.Handler, middlewares ...middlewares.Middleware) {
	h := handler

	for _, m := range middlewares {
		h = m(h)
	}

	http.Handle(pattern, h)
}

func RegisterEndpoints(appContext *AppContext) {
	withDatabaseMiddleware := middlewares.WithDatabase(appContext.DB)

	registerEndpoint("/", http.HandlerFunc(endpoints.HelloEndpoint), middlewares.Logger)
	registerEndpoint("POST /shorten", http.HandlerFunc(endpoints.ShortenEndpoint), middlewares.Logger, withDatabaseMiddleware)
	registerEndpoint("GET /t/{shortenedURLID}", http.HandlerFunc(endpoints.RedirectEndpoint), middlewares.Logger, withDatabaseMiddleware)
}
