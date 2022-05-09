-- +migrate Up
CREATE TABLE tree_samples (
  id         INT(10) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  biid       VARCHAR(26) CHARACTER SET ASCII NOT NULL,
  path       VARCHAR(1350) CHARACTER SET ASCII NOT NULL COMMENT "50 path at maximum",
  office_id  VARCHAR(26) CHARACTER SET ASCII NOT NULL,
  name       VARCHAR(20) NOT NULL,

  valid_from DATE NOT NULL,
  valid_to   DATE NOT NULL DEFAULT "9999-12-31",
  transaction_from DATETIME NOT NULL,
  transaction_to   DATETIME NOT NULL DEFAULT "9999-12-31 23:59:59",

  CONSTRAINT `fk_treesamples_office_id` FOREIGN KEY (`office_id`) REFERENCES `offices` (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS tree_samples;
