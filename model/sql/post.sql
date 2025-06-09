-- post.sql
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`
(
    `id`            int          NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`       int          NOT NULL COMMENT '帖子所属用户id',
    `title`         varchar(255) NOT NULL COMMENT '标题',
    `content`       text         NOT NULL COMMENT '内容',
    `comment_count` int          NOT NULL DEFAULT 0 COMMENT '帖子评论数量',
    `score`         int          NOT NULL DEFAULT 0 COMMENT '帖子分数',
    `type`          int          NOT NULL DEFAULT 0 COMMENT '帖子类型，0正常 1加精',
    `state`         int          NOT NULL DEFAULT 0 COMMENT '帖子状态，0正常 1置顶 2拉黑',
    `create_time`   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    INDEX           `idx_user_id` (`user_id`) -- user_id 索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 添加测试数据
INSERT INTO `post` (user_id, title, content)
VALUES (1, 'Go语言入门教程', '这是一篇关于Go语言的入门文章...'),
       (2, '如何优化MySQL查询性能', '本文介绍了几种优化MySQL查询的方法...'),
       (1, '微服务架构实践', '分享我们团队在微服务转型中的经验和教训...'),
       (3, 'Python数据分析实战', '使用Pandas和Matplotlib进行数据分析的实用技巧...'),
       (2, '前端性能优化指南', '从加载速度到运行效率，全方位提升Web应用性能...'),
       (4, 'Docker容器化部署', '如何使用Docker快速部署和扩展应用...'),
       (1, 'Kubernetes入门到精通', 'Kubernetes核心概念和实际应用案例...'),
       (5, '网络安全基础教程', '了解常见的网络攻击方式和防御策略...'),
       (3, '机器学习基础', '从零开始学习机器学习的数学原理和算法...'),
       (6, '产品经理必备技能', '如何有效管理产品生命周期和团队协作...'),
       (2, '云原生技术发展趋势', '解读云原生技术的未来发展方向...'),
       (4, 'DevOps实践分享', '如何建立高效的DevOps工作流程和文化...'),
       (5, '数据结构与算法', '常见数据结构和算法的Go语言实现...');