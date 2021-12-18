CREATE DATABASE IF NOT EXISTS `gomsx_development` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE USER 'gomsx_user'@'%' IDENTIFIED BY 'gomsx_password';
GRANT ALL PRIVILEGES ON *.* TO 'gomsx_user'@'%' WITH GRANT OPTION;  
