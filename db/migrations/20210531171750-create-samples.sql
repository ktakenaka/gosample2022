-- +migrate Up
CREATE TABLE samples (
  id         INT(10) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  biid       VARCHAR(26) CHARACTER SET ASCII NOT NULL COMMENT "use ulid",
  office_id  VARCHAR(26) CHARACTER SET ASCII NOT NULL,
  code       VARCHAR(10) NOT NULL,
  category   ENUM("small", "medium", "large") NOT NULL,
  amount     DECIMAL(10, 2) UNSIGNED NOT NULL,

  valid_from DATE NOT NULL,
  valid_to   DATE NOT NULL DEFAULT "9999-12-31",
  created_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,

  version TINYINT UNSIGNED NOT NULL,

  CONSTRAINT `fk_samples_office_id` FOREIGN KEY (`office_id`) REFERENCES `offices` (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS samples;
