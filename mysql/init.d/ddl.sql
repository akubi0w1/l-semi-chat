DROP DATABASE `ls_chat`;
CREATE DATABASE IF NOT EXISTS `ls_chat`;

DROP DATABASE `ls_chat`;

CREATE DATABASE IF NOT EXISTS `ls_chat`.`threads_tags`(
    `id` VARCHAR(36) NOT NULL COMMENT 'id',
    `threads_id` VARCHAR(36) NOT NULL COMMENT 'スレッドID',
    `tags_id` VARCHAR(36) NOT NULL COMMENT 'タグID',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_threads_tags`
        FOREIGN KEY (`thread_id`)
        REFERENCES `ls_chat`.`threads` (`thread_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_tags_threads`
        FOREIGN KEY (`tag_id`)
        REFERENCES `ls_chat`.`tags` (`tag_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `unique_thread_tag`
        UNIQUE (`thread_id`,`tag_id`)
)
COMMENT='スレッドのタグ';

CREATE DATABASE IF NOT EXISTS `ls_chat`.`users_tags`(
    `id` VARCHAR(36) NOT NULL COMMENT 'id',
    `users_id` VARCHAR(36) NOT NULL COMMENT 'ユーザーID',
    'tags_id' VARCHAR(36) NOT NULL COMMENT 'タグID',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_users_tags`
        FOREIGN KEY (`user_id`)
        REFERENCES `ls_chat`.`users` (`users_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_tags_users`
        FOREIGN KEY (`tag_id`)
        REFERENCES `ls_chat`.`tags` (`tag_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `unique_user_tag`
        UNIQUE (`user_id`,`tag_id`)
)
COMMENT='ユーザーのタグ';

CREATE DATABASE IF NOT EXISTS `ls_chat`.`users_threads`(
    `id` VARCHAR(36) NOT NULL COMMENT 'id',
    `users_id` VARCHAR(36) NOT NULL COMMENT 'ユーザーID',
    `threads_id` VARCHAR(36) NOT NULL COMMENT 'スレッドID',
    `is_admin` TINYINT NOT NULL DEFAULT 0 COMMENT 'スレッドの管理者判断',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_users_threads`
        FOREIGN KEY (`user_id`)
        REFERENCES `ls_chat`.`users` (`users_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_threads_users`
        FOREIGN KEY (`thread_id`)
        REFERENCES `ls_chat`.`users` (`threads_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `unique_user_thread`
        UNIQUE (`user_id`,`thread_id`)
)
COMMENT='ユーザーのスレッド';

CREATE DATABASE IF NOT EXISTS `ls_chat`.`users_favorites`(
    `id` VARCHAR(36) NOT NULL COMMENT 'id',
    `users_id` VARCHAR(36) NOT NULL COMMENT 'ユーザーID',
    `message_id` VARCHAR(36) NOT NULL COMMENT 'メッセージID',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_users_messages`
        FOREIGN KEY (`user_id`)
        REFERENCES `ls_chat`.`users` (users_id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_messages_users`
        FOREIGN KEY (`message_id`)
        REFERENCES `ls_chat`.`message` (message_id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `unique_user_id_message_id`
        UNIQUE (`user_id`,`message_id`)
)
COMMENT='ユーザーのいいね';

-- define tables
CREATE TABLE IF NOT EXISTS `ls_chat`.`users`(
    `id` VARCHAR(36) PRIMARY KEY COMMENT'id',
    `user_id` VARCHAR(36) UNIQUE NOT NULL COMMENT 'ユーザid',
    `name` VARCHAR(64) NOT NULL COMMENT '名前',
    `image` VARCHAR(128) NOT NULL COMMENT '画像',
    `profile` VARCHAR(150) COMMENT 'プロフィール',
    `is_admin` TINYINT NOT NULL DEFAULT 0 COMMENT '権威',
    `mail` VARCHAR(254) NOT NULL UNIQUE COMMENT 'メールアドレス',
    `login_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'ログイン日時',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `password` VARCHAR(70) NOT NULL COMMENT 'パスワード'
)
COMMENT = 'ユーザ';

CREATE TABLE IF NOT EXISTS `ls_chat`.`messages`(
    `id` VARCHAR(36) NOT NULL COMMENT 'id',
    `message` VARCHAR(150) NOT NULL COMMENT '投稿本文',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `grade` INTEGER UNSIGNED NOT NULL DEFAULT 0 COMMENT '発言のグレード' ,
    `user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
    `thread_id` VARCHAR(36) NOT NULL COMMENT 'スレッドID',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_messages_users`
        FOREIGN KEY (`user_id`)
        REFERENCES `ls_chat`.`users` (`user_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_messages_threads`
        FOREIGN KEY (`thread_id`)
        REFERENCES `ls_chat`.`threads`(`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
)
COMMENT = '投稿メッセージ';

CREATE TABLE IF NOT EXISTS `ls_chat`.`threads`(
    `id` VARCHAR(36) PRIMARY KEY NOT NULL COMMENT 'id',
    `name` VARCHAR(32) NOT NULL COMMENT '名前',
    `description` VARCHAR(150) COMMENT '説明',
    `limit_users` INTEGER COMMENT '上限人数',
    `user_id` VARCHAR(64) NOT NULL COMMENT '管理者',-- F
    `is_public` TINYINT NOT NULL DEFAULT 0 COMMENT '範囲',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
    CONSTRAINT `fk_threads_users`
        FOREIGN KEY (`user_id`)
        REFERENCES `ls_chat`.`users`(`user_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
)
COMMENT ='スレッド';

CREATE TABLE IF NOT EXISTS `ls_chat`.`tags`(
    `id` VARCHAR(36) NOT NULL PRIMARY KEY COMMENT 'id',
    `tag` VARCHAR(25) NOT NULL COMMENT 'タグ名',
    `category_id` VARCHAR(36) NOT NULL COMMENT '大枠',
    CONSTRAINT `fk_tags_categories`
        FOREIGN KEY (`category_id`)
        REFERENCES `ls_chat`.`categories`(`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    UNIQUE(`tag`,`category_id`)
)
COMMENT 'タグ';

CREATE TABLE IF NOT EXISTS `ls_chat`.`categories`(
    `id` VARCHAR(36) PRIMARY KEY NOT NULL COMMENT'id',
    `category` VARCHAR(8) NOT NULL COMMENT '大枠名'
)
COMMENT'カテゴリ';