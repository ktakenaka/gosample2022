# For application
CREATE DATABASE IF NOT EXISTS `gosample2022_development` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE USER 'writer'@'%' IDENTIFIED BY 'writer_password';
GRANT ALL PRIVILEGES ON `gosample2022_development`.* TO 'writer'@'%' WITH GRANT OPTION;  

CREATE USER 'reader' IDENTIFIED BY 'reader_password';
GRANT SELECT, SHOW VIEW ON `gosample2022_development`.* TO 'reader'@'%';

# For testing
CREATE DATABASE IF NOT EXISTS `gosample2022_test1` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `gosample2022_test2` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `gosample2022_test3` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `gosample2022_test4` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `gosample2022_test5` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `gosample2022_test6` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `gosample2022_test7` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `gosample2022_test8` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE USER 'test_writer'@'%' IDENTIFIED BY 'writer_password';
GRANT ALL PRIVILEGES ON `gosample2022_test%`.* TO 'test_writer'@'%' WITH GRANT OPTION;  

CREATE USER 'test_reader' IDENTIFIED BY 'reader_password';
GRANT SELECT, SHOW VIEW ON `gosample2022_test%`.* TO 'test_reader'@'%';
