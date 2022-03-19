# For application
CREATE DATABASE IF NOT EXISTS `gosample2022_development` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE USER 'writer'@'%' IDENTIFIED BY 'writer_password';
GRANT ALL PRIVILEGES ON `gosample2022_development`.* TO 'writer'@'%' WITH GRANT OPTION;  

CREATE USER 'reader' IDENTIFIED BY 'reader_password';
GRANT SELECT, SHOW VIEW ON `gosample2022_development`.* TO 'reader'@'%';

# For testing
CREATE DATABASE IF NOT EXISTS `gosample2022_test` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
