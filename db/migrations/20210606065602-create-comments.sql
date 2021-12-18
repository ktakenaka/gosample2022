
-- +migrate Up
CREATE TABLE sample_comments (
  id         BINARY(16) PRIMARY KEY,
  sample_id  BINARY(16) NOT NULL,
  content    VARCHAR(63),
  CONSTRAINT `fk_sample_comments_sample_id` FOREIGN KEY(sample_id) REFERENCES samples(id),
  INDEX `idx_sample_comments_sample_id` (sample_id)
);

-- +migrate Down
DROP TABLE IF EXISTS sample_comments;
