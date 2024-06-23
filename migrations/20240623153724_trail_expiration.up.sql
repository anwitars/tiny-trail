ALTER TABLE trails
  ADD COLUMN expiration TIMESTAMP NOT NULL
    -- this means that even if our trail was added a year ago, it still gets one day
    DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 day';
