
-- +migrate Up
CREATE TABLE pilots (
  id integer NOT NULL,
  name text NOT NULL
);
ALTER TABLE pilots ADD CONSTRAINT pilot_pkey PRIMARY KEY (id);

-- +migrate Down
DROP TABLE IF EXISTS pilots;
