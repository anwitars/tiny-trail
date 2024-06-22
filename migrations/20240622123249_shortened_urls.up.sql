CREATE TABLE shortened_urls (
  id SERIAL PRIMARY KEY,
  original_url TEXT NOT NULL,
  short_id TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX shortened_urls_short_id_idx ON shortened_urls (short_id);
