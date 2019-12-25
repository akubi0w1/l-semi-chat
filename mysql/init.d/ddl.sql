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
        REFERENCES `ls_chat`.`threads` (`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_tags_threads`
        FOREIGN KEY (`tag_id`)
        REFERENCES `ls_chat`.`tags` (`id`)
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
        REFERENCES `ls_chat`.`users` (`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_tags_users`
        FOREIGN KEY (`tag_id`)
        REFERENCES `ls_chat`.`tags` (`id`)
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
        REFERENCES `ls_chat`.`users` (`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_threads_users`
        FOREIGN KEY (`thread_id`)
        REFERENCES `ls_chat`.`users` (`id`)
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
        REFERENCES `ls_chat`.`users` (`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_messages_users`
        FOREIGN KEY (`message_id`)
        REFERENCES `ls_chat`.`message` (`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `unique_user_id_message_id`
        UNIQUE (`user_id`,`message_id`)
)
COMMENT='ユーザーのいいね';

