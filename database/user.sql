DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`         BIGINT ( 20 ) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`    BIGINT(20) UNSIGNED NOT NULL DEFAULT 0,
    `username`   VARCHAR(64) NOT NULL DEFAULT '',
    `password`   VARCHAR(64) NOT NULL DEFAULT '',
    `email`      VARCHAR(64) NOT NULL DEFAULT '',
    `gender`     TINYINT(4) NOT NULL DEFAULT 0,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_username` (`username`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT="用户表";