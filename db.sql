CREATE TABLE `NewTable` (
`id`  int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键' ,
`username`  varchar(31) NOT NULL COMMENT '用户名' ,
`password`  char(32) NOT NULL COMMENT '密码' ,
`mask`  char(5) NOT NULL COMMENT '掩码' ,
`email`  varchar(63) NOT NULL DEFAULT '' COMMENT '邮件' ,
`created_at`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
`updated_at`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB CHARSET=utf8
COMMENT='用户表';

