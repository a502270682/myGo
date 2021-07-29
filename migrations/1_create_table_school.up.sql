BEGIN;
CREATE TABLE `school` (
                          `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                          `name` varchar(10) DEFAULT NULL COMMENT '名字',
                          `school_type` varchar(10) DEFAULT NULL COMMENT '类型',
                          `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                          `deleted_at` datetime COMMENT '删除时间',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户';
COMMIT;