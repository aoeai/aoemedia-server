CREATE TABLE `file` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `hash` char(64) NOT NULL COMMENT '文件内容的哈希值',
  `size_in_bytes` int unsigned NOT NULL COMMENT '文件大小 单位:字节',
  `filename` varchar(255) NOT NULL COMMENT '文件名',
  `storage_dir` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '存储目录',
  `source` tinyint unsigned NOT NULL COMMENT '来源 1:相机 2:微信',
  `modified_time` datetime NOT NULL COMMENT '修改时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=262 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件';