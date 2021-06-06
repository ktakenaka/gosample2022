
-- +migrate Up
CREATE TABLE sample_comments (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  sample_id  BIGINT UNSIGNED NOT NULL,
  content    VARCHAR(63),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT `fk_sample_comments_sample_id` FOREIGN KEY(sample_id) REFERENCES samples(id),
  INDEX `idx_sample_comments_sample_id` (sample_id)
);

-- +migrate Down
DROP TABLE IF EXISTS sample_comments;
