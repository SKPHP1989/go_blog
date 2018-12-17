CREATE TABLE `bg_admins` (
	`id`  int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键' ,
	`username`  varchar(31) NOT NULL COMMENT '用户名' ,
	`password`  char(32) NOT NULL COMMENT '密码' ,
	`password_mask`  char(5) NOT NULL COMMENT '密码掩码' ,
	`email`  varchar(63) NOT NULL DEFAULT '' COMMENT '邮件' ,
	`last_login_ip` varchar(31) NOT NULL DEFAULT '' COMMENT '最后登录ip地址',
	`created_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
	`updated_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间' ,
	PRIMARY KEY (`id`)
)
ENGINE=InnoDB CHARSET=utf8
COMMENT='管理员表';

CREATE TABLE `bg_article_categories` (
	`id`  int UNSIGNED NOT NULL COMMENT '主键' ,
	`name`  varchar(63) NOT NULL DEFAULT '' COMMENT '标题' ,
	`cover_material_bucket`  varchar(63) NOT NULL DEFAULT '' COMMENT '封面素材桶' ,
	`description`  varchar(127) NOT NULL DEFAULT '' COMMENT '描述' ,
	`status`  tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态(0:未发布,1:已发布,2:下架,3:其他)' ,
	`is_top`  tinyint UNSIGNED NULL DEFAULT 0 COMMENT '是否置顶' ,
	`order_time`  timestamp NOT NULL DEFAULT 0 ON UPDATE CURRENT_TIMESTAMP COMMENT '排序时间' ,
	`created_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
	`updated_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间' ,
	PRIMARY KEY (`id`)
)
ENGINE=InnoDB CHARSET=utf8
COMMENT='文章分类表';

CREATE TABLE `bg_articles` (
	`id`  int UNSIGNED NOT NULL COMMENT '主键' ,
	`title`  varchar(63) NOT NULL DEFAULT '' COMMENT '标题' ,
	`author`  varchar(31) NOT NULL DEFAULT '' COMMENT '作者' ,
	`article_category_id`  int UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类ID' ,
	`cover_material_bucket`  varchar(63) NOT NULL DEFAULT '' COMMENT '封面素材桶' ,
	`description`  varchar(127) NOT NULL DEFAULT '' COMMENT '描述' ,
	`content`  text NOT NULL COMMENT '内容' ,
	`read_num`  int UNSIGNED NOT NULL DEFAULT 0 COMMENT '阅读数量' ,
	`status`  tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态(0:草稿,1:已发布,2:下架,3:其他)' ,
	`is_top`  tinyint UNSIGNED NULL DEFAULT 0 COMMENT '是否置顶' ,
	`admin_id`  int UNSIGNED NULL DEFAULT 0 COMMENT '管理员id' ,
	`order_time`  timestamp NOT NULL DEFAULT 0 ON UPDATE CURRENT_TIMESTAMP COMMENT '排序时间' ,
	`created_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
	`updated_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间' ,
	PRIMARY KEY (`id`)
)
ENGINE=InnoDB CHARSET=utf8
COMMENT='文章表';

CREATE TABLE `bg_materials` (
	`id`  int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键' ,
	`name` varchar(63) NOT NULL DEFAULT '' COMMENT '名字',
	`bucket` varchar(127) NOT NULL DEFAULT '' COMMENT '桶点',
	`admin_id`  int UNSIGNED NULL DEFAULT 0 COMMENT '管理员id' ,
	`created_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
	`updated_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间' ,
	PRIMARY KEY (`id`)
)
ENGINE=InnoDB CHARSET=utf8
COMMENT='素材表';

##初始化数据

set @password_mask = substring(MD5(RAND()),1,5);
INSERT INTO bg_admins (`id` ,`username` ,`password` ,`password_mask`,`email`,`last_login_ip`) VALUES (1,'admin' ,MD5(concat(MD5('123456') ,@password_mask)) ,@password_mask ,'','127.0.0.1');