CREATE TABLE `pb_user` (
	`id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(64) NOT NULL COMMENT 'i.e. forsenlol',
	`twitch_access_token` VARCHAR(64) NULL DEFAULT NULL COMMENT 'User level access-token',
	`twitch_refresh_token` VARCHAR(64) NULL DEFAULT NULL COMMENT 'User level refresh-token',
	PRIMARY KEY (`id`)
)
COMMENT='Users that log in via the web interface'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;
