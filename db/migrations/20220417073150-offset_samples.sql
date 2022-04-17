
-- +migrate Up
CREATE TABLE offset_samples (
  offset INT UNSIGNED NOT NULL PRIMARY KEY
);

-- +migrate Down
DROP TABLE IF EXISTS offset_samples;
