CREATE DATABASE `test` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `test`;

-- drop table `user`; -- 删除表
CREATE TABLE `user` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `uuid` VARCHAR(36) DEFAULT (UUID()),-- UUID()函数返回的UUID的长度为36，其中包含32个字符和4个短横线
    `username` varchar(30) NOT NULL,
    `password` varchar(30) NOT NULL,
    `email` varchar(30) NOT NULL DEFAULT ''
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
desc `user`;
USE `test`;

INSERT INTO `user` (`username`, `password`, `email`) VALUES ('admin', 'admin', ' [email protected] ');
UPDATE `user` SET `email` = '[email protected]' WHERE `username` = 'admin';
INSERT INTO `user` (`username`, `password`, `email`) VALUES ('test', 'test', 'test@gmail.com');
SELECT * FROM `user`;