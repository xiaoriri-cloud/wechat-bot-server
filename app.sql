DROP TABLE IF EXISTS `bot_user_info`;
CREATE TABLE `bot_user_info`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `wid`        varchar(100) DEFAULT '' COMMENT '微信id',
    `account`    varchar(100) DEFAULT '' COMMENT '微信账号',
    `nick_name`  varchar(100) DEFAULT '' COMMENT '昵称',
    `avatar_url` varchar(100) DEFAULT '' COMMENT '头像',
    `nation`     varchar(100) DEFAULT '' COMMENT '国家',
    `province`   varchar(100) DEFAULT '' COMMENT '省份',
    `city`       varchar(100) DEFAULT '' COMMENT '城市',
    `mobile`     varchar(100) DEFAULT '' COMMENT '手机号',
    `gender`     varchar(100) DEFAULT '' COMMENT '性别',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='微信用户';


CREATE TABLE `bot_config`
(
    `name`   varchar(255) NOT NULL,
    `config` text         NOT NULL,
    UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='全局配置表';