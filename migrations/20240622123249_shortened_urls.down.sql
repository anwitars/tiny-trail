DROP TABLE trails CASCADE;

DROP FUNCTION IF EXISTS generate_short_trail_id();

DROP EXTENSION IF EXISTS "pgcrypto";
DROP EXTENSION IF EXISTS "uuid-ossp";
