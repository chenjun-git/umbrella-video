CREATE TABLE `video` (
  `video_id` int NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `title` VARCHAR(256) NOT NULL,
  `store_path` VARCHAR(256) NOT NULL,
  `thumbnail` VARCHAR(256) NOT NULL,
  `description` text NOT NULL,
  `is_charge` tinyint(1) NOT NULL DEFAULT 0,
  `free_time` int NOT NULL DEFAULT 5,
  `theme_id` int NOT NULL,
  `publish_user_id` VARCHAR(32) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='视频主题';