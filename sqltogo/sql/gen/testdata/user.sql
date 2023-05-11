CREATE TABLE `fat_cat` (
                           `TaskID` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户任务自增ID',
                           `UID` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户uid',
                           `MissionPoolID` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '任务池编号',
                           `MissionType` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '任务类型1,2,3,4...',
                           `TaskType` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '任务类型 0是普通任务 1是特殊任务',
                           `State` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT ' 0未解锁;1任务正常执行;2任务已完成;3任务已过期',
                           `TodayDate` date NOT NULL COMMENT '创建日期',
                           `CurrentValue` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '当前完成数,任务参数',
                           `CurrentArg` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT ' 当任务有两个参数时可使用该字段',
                           `EndTime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '今日结束时间',
                           `CreateTime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建记录时间',
                           PRIMARY KEY (`TaskID`),
                           KEY `TodayDate` (`TodayDate`) USING BTREE,
                           KEY `UID` (`UID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8503 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='# 功能：每日任务用户信息表 创建人：张盛钢 创建时间：2022/8/29';