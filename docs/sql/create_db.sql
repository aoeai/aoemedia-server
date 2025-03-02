CREATE TABLE file
(
    id            INT          NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    hash          CHAR(64)     NOT NULL COMMENT '文件内容的哈希值',
    size          INT          NOT NULL COMMENT '文件大小 单位:字节',
    filename      VARCHAR(255) NOT NULL COMMENT '文件名',
    storage_path  VARCHAR(255) NOT NULL COMMENT '存储路径',
    source        TINYINT UNSIGNED NOT NULL COMMENT '来源 1:相机 2:微信',
    modified_time DATETIME     NOT NULL COMMENT '修改时间',
    created_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件';