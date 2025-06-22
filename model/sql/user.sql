-- user.sql
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          int(0) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username`    varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
    `password`    varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
    `email`       varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL COMMENT '邮箱',
    `avatar`      varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像',
    `state`       tinyint(0) NOT NULL Default 0 COMMENT '状态：0正常，1禁用',
    `create_time` timestamp                                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_email` (`email`) -- 添加唯一索引(一条邮箱仅允许注册一条数据)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
