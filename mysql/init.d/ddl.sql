-- archives
CREATE TABLE IF NOT EXISTS `ls_chat`.`archives`(
    `id` VARCHAR(36) PRIMARY KEY COMMENT 'id',
    `path` VARCHAR(128) NOT NULL COMMENT 'ファイルのpath',
    `is_public` TINYINT NOT NULL DEFAULT 1 COMMENT '公開範囲',
    `password` VARCHAR(70) NOT NULL COMMENT 'パスワード' ,
    `thread_id` VARCHAR(36) NOT NULL COMMENT 'スレッドID',
    CONSTRAINT `fk_archives_threads`
        FOREIGN KEY (`thread_id`)
        REFERENCES `ls_chat`.`threads` (`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
)
COMMENT = 'アーカイブ';

-- evaluations(master)
CREATE TABLE IF NOT EXISTS `ls_chat`.`evalutions`(
    `id` VARCHAR(36) PRIMARY KEY COMMENT 'id',
    `item` VARCHAR(10) NOT NULL UNIQUE COMMENT '評価文'
)
COMMENT = '評価';

-- evaluation_scores
CREATE TABLE IF NOT EXISTS `ls_chat`.`evaluation_scores`(
    `id` VARCHAR(36) PRIMARY KEY COMMENT 'id',
    `evaluation_id` VARCHAR(36) NOT NULL COMMENT '評価ID',
    `users_id` VARCHAR(36) NOT NULL COMMENT 'ユーザID',
    `score` INTEGER NOT NULL DEFAULT 0 COMMENT 'スコア' ,
    CONSTRAINT `fk_evaluation_scores_evaluations`
        FOREIGN KEY (`evaluation_id`)
        REFERENCES `ls_chat`.`evaluations` (`evaluation_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_evaluation_scores_users`
        FOREIGN KEY (`user_id`)
        REFERENCES `ls_chat`.`users` (`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `unique_user_evaluation`
        UNIQUE (`user_id`, `evaluation_id`)
)
COMMENT = '評価スコア';

-- users_followers
CREATE TABLE IF NOT EXISTS `ls_chat`.`users_followers`(
    `id` VARCHAR(36) PRIMARY KEY COMMENT 'id',
    `users_id` VARCHAR(36) NOT NULL COMMENT 'ユーザID',
    `followed_user_id` VARCHAR(36) NOT NULL COMMENT 'フォローユーザーID',
    CONSTRAINT `fk_users_followers_users`
        FOREIGN KEY (`user_id`)
        REFERENCES `ls_chat`.`users` (`user_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_users_followers_followed_users`
        FOREIGN KEY (`followed_user_id`)
        REFERENCES `ls_chat`.`followed_users` (`followed_user_id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `unique_user_followed`
        UNIQUE (`user_id`, `followed_id`)
)
COMMENT = 'フォロワー';
	