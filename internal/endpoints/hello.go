package endpoints

import "net/http"

// Basic endpoint that returns "Hello, World!".
func HelloEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
