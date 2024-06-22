package endpoints

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenedURL struct {
	OriginalURL string `json:"original_url"`
	ShortID     string `json:"short_id"`
}

func ShortenEndpoint(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sqlx.DB)
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

	_, err = db.Exec("INSERT INTO shortened_urls (short_id, original_url) VALUES ($1, $2)", shortenedURL.ShortID, shortenedURL.OriginalURL)
	if err != nil {
		http.Error(w, "Error inserting into database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(shortenedURL.ShortID))
}

func generateShortURL() string {
	u := uuid.New()
	hash := sha256.Sum256([]byte(u.String()))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:8]
	return shortURL
}
