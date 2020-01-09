-- create test data
-- users


-- threads
INSERT INTO `ls_chat`.`threads`(`id`,`name`,`description`,`limit_users`,`user_id`,`is_public`,`created_at`,`updated_at`)  VALUES ("11111111-1111-1111-1111-111111111111",”name1”,”description1”,1,”11111111-1111-1111-1111-111111111111”,1,cast('2018-11-24 11:56:40' as datetime),cast('2018-11-24 11:56:40' as datetime));
INSERT INTO `ls_chat`.`threads`(`id`,`name`,`user_id`) VALUES ("22222222-2222-2222-2222-222222222222”,”name1”,”22222222-2222-2222-2222-222222222222”);

-- messages


-- tags


-- categories


-- archives


-- evaluations


-- evaluation_scores


-- users_followers

INSERT INTO `ls_chat`.`users_followers`(`id`,`user_id`,`followed_user_id`) VALUES (“11111111-1111-1111-1111-111111111111”,“11111111-1111-1111-1111-111111111111”,“11111111-1111-1111-1111-111111111111”);

-- users_tags


-- users_threads

INSERT INTO `ls_chat`.`users_threads`(`id`,`user_id`,`thread_id`,`is_admin`) VALUES ("11111111-1111-1111-1111-111111111111","11111111-1111-1111-1111-111111111111","11111111-1111-1111-1111-111111111111",1);
INSERT INTO `ls_chat`.`users_threads`(`id`,`user_id`,`thread_id`) VALUES ("22222222-2222-2222-2222-222222222222","22222222-2222-2222-2222-222222222222","22222222-2222-2222-2222-222222222222");

-- users_favorites


-- threads_tags

INSERT INTO `ls_chat`.`users_followers`(`id`,`thread_id`,`tag_id`) VALUES (“11111111-1111-1111-1111-111111111111”,“11111111-1111-1111-1111-111111111111”,“11111111-1111-1111-1111-111111111111”);


