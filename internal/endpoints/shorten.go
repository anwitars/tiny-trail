package endpoints

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenedURL struct {
	OriginalURL string `json:"original_url"`
	ShortID     string `json:"short_id"`
}

// TODO: database
var shortenedURLs []ShortenedURL

func ShortenEndpoint(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error in request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	shortenedURL := ShortenedURL{
		OriginalURL: req.URL,
		ShortID:     generateShortURL(),
	}

	shortenedURLs = append(shortenedURLs, shortenedURL)

	w.Write([]byte(shortenedURL.ShortID))
}

// TODO: better short URL generation
func generateShortURL() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
