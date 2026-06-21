CREATE DATABASE IF NOT EXISTS jade_grading DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE jade_grading;

DROP TABLE IF EXISTS jade_bracelets;

CREATE TABLE jade_bracelets (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    name VARCHAR(100) NOT NULL COMMENT '手串名称/编号',
    material VARCHAR(50) NOT NULL COMMENT '材质（如：切米蓝、墨碧等）',
    translucency DECIMAL(5,2) NOT NULL COMMENT '透光度 (0-100，越高越好)',
    fineness DECIMAL(5,2) NOT NULL COMMENT '细度 (0-100，越高越好)',
    bead_count INT DEFAULT NULL COMMENT '珠子颗数（选填，NULL表示未填写）',
    score DECIMAL(5,2) DEFAULT NULL COMMENT '计算得出的综合评分',
    grade VARCHAR(20) DEFAULT NULL COMMENT '等级标签（特级、一级、二级、三级）',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id),
    KEY idx_material (material),
    KEY idx_grade (grade)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='和田玉手串质地评定表';
