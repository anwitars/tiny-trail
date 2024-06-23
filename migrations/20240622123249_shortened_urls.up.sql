CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE OR REPLACE FUNCTION generate_short_trail_id() RETURNS TEXT AS $$
DECLARE
  u UUID;
  hash BYTEA;
  short_id TEXT;
BEGIN
  LOOP
    u = uuid_generate_v4();
    hash = digest(u::TEXT, 'sha256');
    short_id = encode(hash, 'base64')::TEXT;
    short_id = replace(short_id, '/', '_');
    short_id = replace(short_id, '+', '-');
    short_id = replace(short_id, '=', '');
    short_id = left(short_id, 8);

    IF NOT EXISTS (SELECT 1 FROM trails WHERE id = short_id) THEN
      RETURN short_id;
    END IF;
  END LOOP;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE trails (
  id TEXT NOT NULL PRIMARY KEY DEFAULT generate_short_trail_id(),
  original_url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
