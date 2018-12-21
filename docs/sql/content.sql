-- pf内容表 https://segmentfault.com/q/1010000003856705
CREATE TABLE IF NOT EXISTS pf_contents
(
  id        BIGINT UNSIGNED AUTO_INCREMENT  NOT NULL COMMENT 'id',
  content_id   BIGINT UNSIGNED NOT NULL COMMENT 'content唯一标识'，
  title VARCHAR(60) NOT NULL COMMENT 'content的标题',
  description VARCHAR(300) NOT NULL COMMENT 'content的简介或概要',
  author_id BIGINT UNSIGNED NOT NULL COMMENT '作者的user_id',
  category VARCHAR(60) COMMENT '分类，可空'
  type TINY UNSIGNED NOT NULL COMMENT '内容类型，枚举。{0: text, 1: html, 2: markdown}',
  body TEXT NOT NULL COMMENT '内容'，
  version INT UNSIGNED NOT NULL COMMENT '版本，对应审计表中的版本', 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   NOT NULL COMMENT '内容创建时间',
  updated_at TIMESTAMP CURRENT_TIMESTAMP COMMENT '内容更新时间',
  deleted_at  TIMESTAMP CURRENT_TIMESTAMP COMMENT '内容删除时间，软删除标识，实现记住用View',
  extra  TEXT COMMENT '扩展部分，json字符串，不支持索引',
  PRIMARY KEY (id)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_bin
  COMMENT ='content表';

ALTER TABLE pf_contents ADD CONSTRAINT pf_contents_uk UNIQUE (content_id);

-- 内容历史修改表 audit table,
CREATE TABLE IF NOT EXISTS pf_contents_audit
(
  `id`         BIGINT UNSIGNED                          NOT NULL COMMENT 'id',
  `content_id`    BIGINT UNSIGNED                          NOT NULL COMMENT '内容id',
  `version` INT UNSEIGNED             NOT NULL COMMENT '版本，从1开始',
  `json_diff`  TEXT                   NOT NULL COMMENT 'jsondiff(text)',
  `new_value`  TEXT                   NOT NULL COMMENT '新值(text)',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP   NOT NULL COMMENT '内容记录添加时间',
  PRIMARY KEY (id)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_bin
  COMMENT ='内容历史审计表';