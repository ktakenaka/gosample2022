
-- +migrate Up
CREATE TABLE pilots (
  id integer NOT NULL,
  name text NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS pilots;
