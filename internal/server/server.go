package server

import (
	"net/http"
	"tinytrail/internal/endpoints"
	"tinytrail/internal/middlewares"
)

func registerEndpoint(pattern string, handler http.Handler, middlewares ...middlewares.Middleware) {
	h := handler

	for _, m := range middlewares {
		h = m(h)
	}

	http.Handle(pattern, h)
}

func RegisterEndpoints() {
	registerEndpoint("/", http.HandlerFunc(endpoints.HelloEndpoint), middlewares.Logger)
	registerEndpoint("POST /shorten", http.HandlerFunc(endpoints.ShortenEndpoint), middlewares.Logger)
}
