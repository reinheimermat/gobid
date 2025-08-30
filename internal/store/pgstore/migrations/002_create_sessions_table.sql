-- Write your migrate up statements here
CREATE TABLE sessions (
  token TEXT PRIMARY KEY,
  data BYTEA NOT NULL,
  expiry TIMESTAMPTZ NOT NULL
)

---- create above / drop below ----

DROP INDEX IF EXISTS sessions_expiry_idx;
DROP TABLE IF EXISTS sessions;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
