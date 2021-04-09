DROP TABLE IF EXISTS `t_gateway_info`;

CREATE TABLE `t_gateway_info` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `ip` varchar(32) NOT NULL,
    `name` varchar(32) NOT NULL,
    `area` varchar(128) NOT NULL,
    `status` tinyint(1) unsigned NOT NULL DEFAULT 0,
    `update_time` BIGINT(20)   NOT NULL DEFAULT '0' COMMENT '更新时间',
    `create_time` BIGINT(20)   NOT NULL DEFAULT '0' COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `ani` (`area`,`name`,`ip`),
    KEY (`name`),
    KEY (`ip`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT ='网关表';

DROP TABLE IF EXISTS `t_gateway_heart`;

CREATE TABLE `t_gateway_heart` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `gateway_id` int(11) unsigned NOT NULL,
    `status` tinyint(1) unsigned NOT NULL DEFAULT 0,
    `update_time` BIGINT(20)   NOT NULL DEFAULT '0' COMMENT '更新时间',
    `create_time` BIGINT(20)   NOT NULL DEFAULT '0' COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `gc` (`gateway_id`,`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT ='网关心跳表';