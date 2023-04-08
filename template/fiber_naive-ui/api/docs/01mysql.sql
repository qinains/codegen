CREATE TABLE
    `casbin_rule` (
        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '权限ID',
        `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'PType',
        `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'v0',
        `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'v1',
        `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'v2',
        `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'v3',
        `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'v4',
        `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'v5',
        PRIMARY KEY (`id`) USING BTREE,
        UNIQUE KEY `unique_index` (
            `ptype`,
            `v0`,
            `v1`,
            `v2`,
            `v3`,
            `v4`,
            `v5`
        ) USING BTREE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '权限' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `config` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '配置ID',
        `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
        `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '内容',
        `type` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '类型:input,textarea,select,radio,checkbox,datetimepicker,richtext',
        `label` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签',
        `placeholder` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '占位文本',
        `unit` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '单位',
        `extra` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '额外信息，json',
        `sort` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序，值越大越靠前',
        PRIMARY KEY (`id`) USING BTREE,
        INDEX `name`(`name` ASC) USING BTREE
    ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '配置' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `department` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '部门ID',
        `tenant_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '租户ID',
        `parent_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '父部门ID',
        `sort` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序：越大越靠前',
        `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
        `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态:10启用、20禁用',
        `create_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
        `update_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
        `update_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
        `delete_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
        `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
        PRIMARY KEY (`id`) USING BTREE,
        INDEX `parent_id`(`parent_id` ASC) USING BTREE
    ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '部门' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `dict` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '字典ID',
        `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称,英文无空格',
        `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
        `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '类型',
        `description` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
        `sort` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序，越大越靠前',
        `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态:10启用，20禁用',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间，毫秒时间戳',
        `update_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间，毫秒时间戳',
        PRIMARY KEY (`id`) USING BTREE,
        UNIQUE INDEX `name`(`name` ASC) USING BTREE
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '字典' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `dict_item` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '字典项ID',
        `dict_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '字典ID',
        `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称,英文无空格',
        `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
        `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '值',
        `description` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
        `is_default` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否是默认：10是，20否',
        `default` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '默认值',
        `tag_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'tag类型',
        `sort` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序，越大越靠前',
        `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态:10启用，20禁用',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间，毫秒时间戳',
        `update_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间，毫秒时间戳',
        PRIMARY KEY (`id`) USING BTREE,
        INDEX `dict_id`(`dict_id` ASC) USING BTREE
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '字典项' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `file` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文件ID',
        `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
        `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '上传的文件名',
        `url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件url',
        `width` smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT '宽,如果是图片该项有值',
        `height` smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT '高,如果是图片该项有值',
        `size` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小,字节',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间，毫秒时间戳',
        PRIMARY KEY (`id`) USING BTREE
    ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '文件' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `job` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
        `tenant_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '租户ID',
        `parent_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '父岗位ID',
        `sort` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序：越大越靠前',
        `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '岗位名称',
        `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态:10启用、20禁用',
        `create_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
        `update_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
        `update_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
        `delete_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
        `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
        PRIMARY KEY (`id`) USING BTREE,
        INDEX `parent_id`(`parent_id` ASC) USING BTREE
    ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '岗位' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `log_login` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '登录日志ID',
        `login_status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录状态：10成功，20登录失败',
        `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作者ID',
        `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作者',
        `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
        `address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP所在省市区',
        `browser` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '浏览器',
        `system` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作系统',
        `user_agent` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '浏览器头',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间，毫秒时间戳',
        PRIMARY KEY (`id`) USING BTREE
    ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '登录日志' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `log_operation` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '操作日志ID',
        `content` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作内容',
        `in` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作参数',
        `out` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作结果',
        `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作者ID',
        `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作者',
        `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
        `address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP所在省市区',
        `browser` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '浏览器',
        `system` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作系统',
        `user_agent` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '浏览器头',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间，毫秒时间戳',
        PRIMARY KEY (`id`) USING BTREE
    ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '操作日志' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `menu` (
        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
        `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父菜单ID',
        `sort` bigint unsigned NOT NULL DEFAULT '0' COMMENT '排序：越大越靠前',
        `type` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型：10目录、20菜单、30操作',
        `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
        `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
        `permission` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限标识',
        `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'vue路径',
        `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'vue组件',
        `alias` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'vue别名',
        `props` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'vue路由组件传参',
        `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'vue重定向',
        `query` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'vue查询',
        `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标',
        `target` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '打开方式',
        `active_menu` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '激活菜单',
        `affix` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否固定多页签',
        `is_always_show` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否总是显示',
        `is_cache` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否是缓存',
        `is_root` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否是根路由',
        `is_show` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否显示',
        `is_external` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否是外链',
        `url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '外链URL',
        `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态:10启用、20禁用',
        `create_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间，毫秒时间戳',
        `update_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间，毫秒时间戳',
        PRIMARY KEY (`id`) USING BTREE,
        KEY `parent_id` (`parent_id`) USING BTREE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '菜单' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `role` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
        `tenant_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '租户ID',
        `parent_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '父角色ID',
        `sort` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序：越大越靠前',
        `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名',
        `is_default` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否是默认：10是，20否',
        `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
        `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态:10启用、20禁用',
        `create_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
        `update_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
        `update_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
        `delete_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
        `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
        PRIMARY KEY (`id`) USING BTREE,
        INDEX `parent_id`(`parent_id` ASC) USING BTREE
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '角色' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `tenant` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '租户ID',
        `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '编码',
        `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
        `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '描述',
        `start_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '开始时间，毫秒时间戳',
        `end_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '结束时间，毫秒时间戳',
        `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态:10启用、20禁用',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间，毫秒时间戳',
        `update_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间，毫秒时间戳',
        PRIMARY KEY (`id`) USING BTREE,
        INDEX `name`(`name` ASC) USING BTREE
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '租户' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `user` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
        `tenant_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '租户ID',
        `login_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录名',
        `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
        `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号码',
        `password_hash` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
        `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态:10启用、20禁用',
        `create_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
        `update_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
        `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
        `update_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
        `delete_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
        `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
        PRIMARY KEY (`id`) USING BTREE,
        UNIQUE INDEX `login_name`(
            `login_name` ASC,
            `tenant_id` ASC
        ) USING BTREE,
        UNIQUE INDEX `phone`(`phone` ASC, `tenant_id` ASC) USING BTREE
    ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户' ROW_FORMAT = DYNAMIC;

CREATE TABLE
    `user_role` (
        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户角色关联ID',
        `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
        `role_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
        PRIMARY KEY (`id`) USING BTREE,
        UNIQUE INDEX `user_role`(`user_id` ASC, `role_id` ASC) USING BTREE
    ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户角色关联' ROW_FORMAT = DYNAMIC;