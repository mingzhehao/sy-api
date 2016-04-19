DROP TABLE IF EXISTS app;

CREATE TABLE `app` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `app_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
      `app_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
      `create_date` int(11) NOT NULL,
      `publish_date` int(11) DEFAULT NULL,
      PRIMARY KEY (`id`),
      UNIQUE KEY `app_code` (`app_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='APP信息表'
