
-- +migrate Up
CREATE TABLE samples (
  id         BINARY(16) PRIMARY KEY,
  office_id  BINARY(16) NOT NULL,
  title      VARCHAR(30) NOT NULL,
  category   ENUM("small", "medium", "large") NOT NULL,
  memo       TEXT NOT NULL,
  date       DATE NOT NULL,
  amount     DECIMAL(10, 2) UNSIGNED NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT `fk_samples_office_id` FOREIGN KEY (`office_id`) REFERENCES `offices` (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS samples;
