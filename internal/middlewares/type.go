package middlewares

import "net/http"

// Middleware is a function that takes a http.Handler and returns a http.Handler.
// This is used to wrap a http.Handler with additional functionality.
type Middleware func(http.Handler) http.Handler
