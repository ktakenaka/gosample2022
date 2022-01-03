
-- +migrate Up
CREATE TABLE samples (
  id         BINARY(16) PRIMARY KEY,
  office_id  BINARY(16) NOT NULL,
  title      VARCHAR(30) NOT NULL,
  category   ENUM("small", "medium", "large") NOT NULL,
  memo       TEXT NOT NULL,
  date       DATE NOT NULL,
  amount     DECIMAL(10, 2) NOT NULL,
  created_by BINARY(16) NOT NULL,
  updated_by BINARY(16) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `index_samples_office_id`(`office_id`),
  CONSTRAINT `fk_samples_office_id` FOREIGN KEY (`office_id`) REFERENCES `offices` (`id`),
  CONSTRAINT `fk_samples_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_samples_updated_by` FOREIGN KEY (`updated_by`) REFERENCES `users` (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS samples;
