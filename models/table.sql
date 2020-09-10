CREATE  Table `user` (
    `id` bigint(20) NOT NUll AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL ,
    `username` varchar(64) collate utf8mb4_general_ci NOT NULL,
    `password` varchar(64) collate utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT 0,
    `create_time` timestamp NULL default current_timestamp,
    `udppate_time` timestamp NULL default current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE ,
    UNIQUE KEY `idx_uer_id` (`user_id`) USING BTREE

) engine = InnoDB DEFAULT CHAR SET=utf8mb4 collate=utf8mb4_general_ci ;


SELECT id,



DROP  TABLE  IF EXISTS `coumunity`;
CREATE TABLE `coumunity` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `community_id` int(10) unsigned NOT NULL,
    `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL ,
    `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `create_time` timestamp NULL default current_timestamp,
    `udpate_time` timestamp NULL default current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_username` (`community_name`) USING BTREE ,
    UNIQUE KEY `idx_community_id` (`community_id`) USING BTREE
)engine = InnoDB DEFAULT CHAR SET=utf8mb4 collate=utf8mb4_general_ci ;

INSERT INTO  `coumunity` VALUES ('1','1','radial','golanging,go,go','2020-08-11 21:34:47','2020-08-11 21:34:47');
INSERT INTO  `coumunity` VALUES ('2','2','python','import this','2020-08-11 21:34:47','2020-08-12 21:34:47');
INSERT INTO  `coumunity` VALUES ('3','3','react','babel','2020-08-11 20:34:47','2020-08-12 21:34:47');
INSERT INTO  `coumunity` VALUES ('4','5','golang','def main() {}','2020-08-21 21:34:47','2020-09-11 21:34:47');

# SELECT community_id,community_name,introduction,create_time,udpate_time FROM coumunity WHERE  community_id = ?

DROP  TABLE  IF EXISTS `post`;
CREATE TABLE `post` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(10) unsigned NOT NULL,
    `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL ,
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL ,
    `author_id` bigint(20) NOT NULL ,
    `community_id` bigint(28) NOT NULL ,
    `status` tinyint(4) NOT NULL DEFAULT '1',
    `create_time` timestamp NULL default current_timestamp,
    `update_time` timestamp NULL default current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_community_id` (`community_id`),
    KEY `idx_community_id` (`community_id`)
)engine = InnoDB DEFAULT CHAR SET=utf8mb4 collate=utf8mb4_general_ci ;


select
       post_id,
       title,
       content,
       author_id,
       community_id,
       status,
       create_time,
       update_time
       from post WHERE post_id = ?

