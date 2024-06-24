package middlewares

import "net/http"

func Apply(h http.Handler, m ...Middleware) http.Handler {
	for _, middleware := range m {
		h = middleware(h)
	}
	return h
}
