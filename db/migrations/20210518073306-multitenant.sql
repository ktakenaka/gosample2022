-- +migrate Up
CREATE TABLE IF NOT EXISTS  `offices`(
    `id`   VARCHAR(26) CHARACTER SET ASCII PRIMARY KEY,
    `name` VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS `users`(
    `id`    VARCHAR(26) CHARACTER SET ASCII PRIMARY KEY,
    `email` VARCHAR(100) NOT NULL,
    UNIQUE KEY `uk_users_email`(`email`)
);

CREATE TABLE IF NOT EXISTS `office_users`(
    `id`        VARCHAR(26) CHARACTER SET ASCII PRIMARY KEY,
    `office_id` VARCHAR(26) CHARACTER SET ASCII NOT NULL,
    `user_id`   VARCHAR(26) CHARACTER SET ASCII NOT NULL,
    CONSTRAINT `fk_office_users_office_id` FOREIGN KEY (`office_id`) REFERENCES `offices`(`id`),
    CONSTRAINT `fk_office_users_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    CONSTRAINT UNIQUE KEY `uk_office_user`(`office_id`, `user_id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `office_users`;
DROP TABLE IF EXISTS `offices`;
DROP TABLE IF EXISTS `users`;
