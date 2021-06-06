
-- +migrate Up
CREATE TABLE samples (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  title      VARCHAR(20) NOT NULL,
  category   ENUM("small", "medium", "large"),
  memo       TEXT DEFAULT NULL,
  date       DATE NOT NULL,
  amount     DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS samples;
