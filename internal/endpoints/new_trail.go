package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"
	"tinytrail/internal/types"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Input JSON for the shorten endpoint.
type ShortenRequest struct {
	URL        string               `json:"url"`
	Expiration types.Nullable[uint] `json:"expiration"` // in hours
}

// Output JSON for the shorten endpoint.
type ShortenedURL struct {
	OriginalURL string    `json:"original_url"`
	ShortID     string    `json:"short_id"`
	Expiration  time.Time `json:"expiration"` // in hours
}

func (s *ShortenRequest) Validate() error {
	parsedURL, err := url.ParseRequestURI(s.URL)
	if err != nil {
		return errors.New("invalid URL")
	}

	slog.Info(fmt.Sprintf("Parsed URL: %+v", parsedURL))

	if s.Expiration.Valid && s.Expiration.Value <= 0 {
		return errors.New("expiration must be greater than 0")
	}

	return nil
}

// Shortens a URL and returns a short URL.
func NewTrailEndpoint(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sqlx.DB)
	var req ShortenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	shortenedURL := ShortenedURL{
		OriginalURL: req.URL,
	}

	var insert_statement string
	var insert_args []interface{}

	// TODO: should refactor this to easily add more fields in the future
	if !req.Expiration.Valid {
		insert_statement = "INSERT INTO trails (original_url) VALUES ($1) RETURNING id"
		insert_args = []interface{}{shortenedURL.OriginalURL}
	} else {
		insert_statement = "INSERT INTO trails (original_url, expiration) VALUES ($1, (CURRENT_TIMESTAMP + CONCAT($2::TEXT, ' hours')::interval)) RETURNING id"
		insert_args = []interface{}{shortenedURL.OriginalURL, req.Expiration.Value}
	}

	err = db.QueryRowx(insert_statement, insert_args...).Scan(&shortenedURL.ShortID)
	if err != nil {
		http.Error(w, "Error inserting into database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(shortenedURL.ShortID))
}
