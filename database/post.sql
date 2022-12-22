DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`
(
    `id`           BIGINT ( 20 ) UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id`      BIGINT ( 20 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子id',
    `user_id`      BIGINT ( 20 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '作者的用户id',
    `community_id` BIGINT ( 20 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属社区',
    `title`        VARCHAR(128) NOT NULL DEFAULT '' COMMENT '标题',
    `content`      TEXT COMMENT '内容',
    `status`       TINYINT ( 4 ) NOT NULL DEFAULT '1' COMMENT '帖子状态',
    `created_at`   TIMESTAMP NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`   TIMESTAMP NULL DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY            `idx_post_id` ( `post_id` ),
    KEY            `idx_user_id` ( `user_id` ),
    KEY            `idx_community_id` ( `community_id` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;