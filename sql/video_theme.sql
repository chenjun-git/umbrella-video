CREATE TABLE `video_theme` (
  `theme_id` 	int 			NOT NULL AUTO_INCREMENT,
  `name` 		VARCHAR(128) 	NOT NULL DEFAULT '',
  `description` text,
  `created_at` 	timestamp 		NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` 	timestamp 		NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`theme_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;