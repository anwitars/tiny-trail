package endpoints

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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

func generateShortURL() string {
	u := uuid.New()
	hash := sha256.Sum256([]byte(u.String()))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:8]
	return shortURL
}
