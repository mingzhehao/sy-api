CREATE DATABASE sy_vendor
  DEFAULT CHARACTER SET utf8
  DEFAULT COLLATE utf8_general_ci;
USE sy_vendor;
SET NAMES utf8;

DROP TABLE IF EXISTS user;

CREATE TABLE `user` (
      `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户编号',
      `user_name` varchar(45) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名称',
      `user_age` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户年龄',
      `user_sex` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
      PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户表' 
