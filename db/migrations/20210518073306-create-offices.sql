-- +migrate Up
CREATE TABLE IF NOT EXISTS  `offices`(
    `id`   BINARY(16) PRIMARY KEY,
    `name` VARCHAR(30) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS `offices`;
