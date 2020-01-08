-- create test data
-- users
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`image`,`profile`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("11111111-1111-1111-1111-111111111111","userID1","name1","hoge@example.com","/image/path","profile1",1,cast('2019/10/11 08:08:08' as datetime),cast('2019/10/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`profile`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("22222222-2222-2222-2222-222222222222","userID2","name1","fuga@example.com","profile1",0,cast('2019/09/11 08:08:08' as datetime),cast('2010/10/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`image`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("33333333-3333-3333-3333-333333333333","userID3","name3","hoge@sample.com","/image/path",0,cast('2018/06/11 08:08:08' as datetime),cast('2011/08/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`image`,`profile`,`login_at`,`created_at`,`password`) VALUES ("44444444-4444-4444-4444-444444444444","userID4","name4","fuga@sample.com","/image/path","profile",cast('2018/03/11 08:08:08' as datetime),cast('2016/10/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`image`,`profile`,`is_admin`,`created_at`,`password`) VALUES ("55555555-5555-5555-5555-555555555555","userID5","name5","hoge@hogeample.com","/image/path","profile",0,cast('2019/05/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`image`,`profile`,`is_admin`,`login_at`,`password`) VALUES ("66666666-6666-6666-6666-666666666666","userID6","name6","fuga@hogeample.com","/image/path","profile",0,cast('2019/10/11 08:08:07' as datetime),"password1");
---- 異常系(id null, user_id null, name null, mail null,password null, id prim, userId prim, mail prim)
INSERT INTO `ls_chat`.`users` (`user_id`,`name`,`mail`,`image`,`profile`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("userID7","name7","hoge@fugample.com","/image/path","profile",0,cast('2019/12/25 08:08:08' as datetime),cast('2019/10/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`name`,`mail`,`image`,`profile`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("77777777-7777-7777-7777-777777777777","name7","hoge@fugample.com","/image/path","profile",0,cast('2019/10/11 08:08:08' as datetime),cast('2019/10/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`mail`,`image`,`profile`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("77777777-7777-7777-7777-777777777777","userID7","hoge@fugample.com","/image/path","profile",0,cast('2019/10/11 08:08:08' as datetime),cast('2019/10/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`image`,`profile`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("77777777-7777-7777-7777-777777777777","userID7","name7","/image/path","profile",0,cast('2019/10/11 08:08:08' as datetime),cast('2019/10/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`image`,`profile`,`is_admin`,`login_at`,`created_at`) VALUES ("77777777-7777-7777-7777-777777777777","userID7","name7","hoge@fugample.com","/image/path","profile",0,cast('2019/10/11 08:08:08' as datetime),cast('2019/10/11 08:08:07' as datetime));
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`image`,`profile`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("11111111-1111-1111-1111-111111111111","userID7","name7","hoge@fugample.com","/image/path","profile",0,cast('2019/10/11 08:08:08' as datetime),cast('2019/10/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`image`,`profile`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("77777777-7777-7777-7777-777777777777","userID1","name7","hoge@fugample.com","/image/path","profile",0,cast('2019/10/11 08:08:08' as datetime),cast('2019/10/11 08:08:07' as datetime),"password1");
INSERT INTO `ls_chat`.`users` (`id`,`user_id`,`name`,`mail`,`image`,`profile`,`is_admin`,`login_at`,`created_at`,`password`) VALUES ("77777777-7777-7777-7777-777777777777","userID7","name7","hoge@example.com","/image/path","profile",0,cast('2019/10/11 08:08:08' as datetime),cast('2019/10/11 08:08:07' as datetime),"password1");


-- threads


-- messages
INSERT INTO `ls_chat`.`messages` (id,message,created_at,grede,user_id,thread_id) VALUES ('11111111-1111-1111-1111-111111111111','おはよう','2019-12-09 07:14:22',DEFAULT,'11111111-1111-1111-1111-111111111111','11111111-1111-1111-1111-111111111111');
INSERT INTO `ls_chat`.`messages` (id,message,created_at,grede,user_id,thread_id) VALUES ('22222222-2222-2222-2222-222222222222','こんにちは','2019-12-10 11:24:30',DEFAULT,'22222222-2222-2222-2222-222222222222','22222222-2222-2222-2222-222222222222');
INSERT INTO `ls_chat`.`messages` (id,message,created_at,grede,user_id,thread_id) VALUES ('33333333-3333-3333-3333-333333333333','こんばんわ','2019-12-11 19:56:46',DEFAULT,'33333333-3333-3333-3333-333333333333','33333333-3333-3333-3333-333333333333');
INSERT INTO `ls_chat`.`messages` (id,message,created_at,grede,user_id,thread_id) VALUES ('44444444-4444-4444-4444-444444444444','犬こそが至高',DEFAULT,1,'44444444-4444-4444-4444-444444444444','44444444-4444-4444-4444-444444444444');
INSERT INTO `ls_chat`.`messages` (id,message,created_at,grede,user_id,thread_id) VALUES ('55555555-5555-5555-5555-555555555555','猫でしょ',DEFAULT,2,'55555555-5555-5555-5555-555555555555','55555555-5555-5555-5555-555555555555');
INSERT INTO `ls_chat`.`messages` (id,message,created_at,grede,user_id,thread_id) VALUES ('66666666-6666-6666-6666-666666666666','(鳥なんだよなあ)',DEFAULT,3,'66666666-6666-6666-6666-666666666666','66666666-6666-6666-6666-666666666666');
-- tags
INSERT INTO `ls_chat`.`tags`(`id`,`tag`,`category_id`) VALUES ("11111111-1111-1111-1111-111111111111","tag1","11111111-1111-1111-1111-111111111111");
---- 異常系(id null, tag null, category null, id prim, unique)
INSERT INTO `ls_chat`.`tags`(`tag`,`category_id`) VALUES ("tag2","11111111-1111-1111-1111-111111111111");
INSERT INTO `ls_chat`.`tags`(`id`,`category_id`) VALUES ("22222222-2222-2222-2222-222222222222","11111111-1111-1111-1111-111111111111");
INSERT INTO `ls_chat`.`tags`(`id`,`tag`) VALUES ("22222222-2222-2222-2222-222222222222","tag2");
INSERT INTO `ls_chat`.`tags`(`id`,`tag`,`category_id`) VALUES ("11111111-1111-1111-1111-111111111111","tag2","11111111-1111-1111-1111-111111111111");
INSERT INTO `ls_chat`.`tags`(`id`,`tag`,`category_id`) VALUES ("22222222-2222-2222-2222-222222222222","tag1","11111111-1111-1111-1111-111111111111");


-- categories
INSERT INTO `ls_chat`.`categories`(`id`,`category`) VALUES ("11111111-1111-1111-1111-111111111111","category1");
---- 異常系(id null, category null, id prim, cate uni?)
INSERT INTO `ls_chat`.`categories`(`id`,`category`) VALUES ("category2");
INSERT INTO `ls_chat`.`categories`(`id`) VALUES ("22222222-2222-2222-2222-222222222222");
INSERT INTO `ls_chat`.`categories`(`id`,`category`) VALUES ("11111111-1111-1111-1111-111111111111","category2");
INSERT INTO `ls_chat`.`categories`(`id`,`category`) VALUES ("22222222-2222-2222-2222-222222222222","category1");


-- archives
INSERT INTO `ls_chat`.`archives` (id,path,is_public,password,thread_id) VALUES ('11111111-1111-1111-1111-111111111111','./path',DEFAULT,'password','11111111-1111-1111-1111-111111111111');
INSERT INTO `ls_chat`.`archives` (id,path,is_public,password,thread_id) VALUES ('22222222-2222-2222-2222-222222222222','./User1',DEFAULT,'12345678','22222222-2222-2222-2222-222222222222');
INSERT INTO `ls_chat`.`archives` (id,path,is_public,password,thread_id) VALUES ('33333333-3333-3333-3333-333333333333','./User2',1,'qwertyui','33333333-3333-3333-3333-333333333333');
INSERT INTO `ls_chat`.`archives` (id,path,is_public,password,thread_id) VALUES ('44444444-4444-4444-4444-444444444444','./User3',1,'asdfghjk','44444444-4444-4444-4444-444444444444');
INSERT INTO `ls_chat`.`archives` (id,path,is_public,password,thread_id) VALUES ('55555555-5555-5555-5555-555555555555','./User4',0,'AsDfGhJk','55555555-5555-5555-5555-555555555555');
INSERT INTO `ls_chat`.`archives` (id,path,is_public,password,thread_id) VALUES ('66666666-6666-6666-6666-666666666666','./User5',0,'qawsedrftgyhujikolp','66666666-6666-6666-6666-666666666666');
-- evaluations
INSERT INTO `ls_chat`.`evaluations`(`id`,`item`) VALUES ("11111111-1111-1111-1111-111111111111","item1");
---- 異常系(prim, item not unique, item is null)
INSERT INTO `ls_chat`.`evaluations`(`id`,`item`) VALUES ("11111111-1111-1111-1111-111111111111","item2");
INSERT INTO `ls_chat`.`evaluations`(`id`,`item`) VALUES ("22222222-2222-2222-2222-222222222222","item1");
INSERT INTO `ls_chat`.`evaluations`(`id`) VALUES ("33333333-3333-3333-3333-333333333333");


-- evaluation_scores
INSERT INTO `ls_chat`.`categories`(`id`,`evaluation_id`,`user_id`,`score`) VALUES ("11111111-1111-1111-1111-111111111111","11111111-1111-1111-1111-111111111111","11111111-1111-1111-1111-111111111111",0);
INSERT INTO `ls_chat`.`categories`(`id`,`evaluation_id`,`user_id`) VALUES ("22222222-2222-2222-2222-222222222222","11111111-1111-1111-1111-111111111111","22222222-2222-2222-2222-222222222222");
---- 異常系(id null, evalu null, user null, id prim, unique)
INSERT INTO `ls_chat`.`categories`(`evaluation_id`,`user_id`,`score`) VALUES ("11111111-1111-1111-1111-111111111111","33333333-3333-3333-3333-333333333333",0);
INSERT INTO `ls_chat`.`categories`(`id`,`user_id`,`score`) VALUES ("33333333-3333-3333-3333-333333333333","33333333-3333-3333-3333-333333333333",0);
INSERT INTO `ls_chat`.`categories`(`id`,`evaluation_id`,`score`) VALUES ("33333333-3333-3333-3333-333333333333","11111111-1111-1111-1111-111111111111",0);
INSERT INTO `ls_chat`.`categories`(`id`,`evaluation_id`,`user_id`,`score`) VALUES ("11111111-1111-1111-1111-111111111111","11111111-1111-1111-1111-111111111111","33333333-3333-3333-3333-333333333333",0);
INSERT INTO `ls_chat`.`categories`(`id`,`evaluation_id`,`user_id`,`score`) VALUES ("33333333-3333-3333-3333-333333333333","11111111-1111-1111-1111-111111111111","11111111-1111-1111-1111-111111111111",0);


-- users_followers


-- users_tags
INSERT INTO `ls_chat`.`users_tags` (id,user_id,tag_id) VALUES ('11111111-1111-1111-1111-111111111111','11111111-1111-1111-1111-111111111111','11111111-1111-1111-1111-111111111111');
INSERT INTO `ls_chat`.`users_tags` (id,user_id,tag_id) VALUES ('22222222-2222-2222-2222-222222222222','22222222-2222-2222-2222-222222222222','22222222-2222-2222-2222-222222222222');
INSERT INTO `ls_chat`.`users_tags` (id,user_id,tag_id) VALUES ('33333333-3333-3333-3333-333333333333','33333333-3333-3333-3333-333333333333','33333333-3333-3333-3333-333333333333');
INSERT INTO `ls_chat`.`users_tags` (id,user_id,tag_id) VALUES ('44444444-4444-4444-4444-444444444444','44444444-4444-4444-4444-444444444444','44444444-4444-4444-4444-444444444444');
INSERT INTO `ls_chat`.`users_tags` (id,user_id,tag_id) VALUES ('55555555-5555-5555-5555-555555555555','55555555-5555-5555-5555-555555555555','55555555-5555-5555-5555-555555555555');
INSERT INTO `ls_chat`.`users_tags` (id,user_id,tag_id) VALUES ('66666666-6666-6666-6666-666666666666','66666666-6666-6666-6666-666666666666','66666666-6666-6666-6666-666666666666');
-- users_threads


-- users_favorites
INSERT INTO `ls_chat`.`users_favorites` (id,user_id,message_id) VALUES ('11111111-1111-1111-1111-111111111111','11111111-1111-1111-1111-111111111111','11111111-1111-1111-1111-111111111111');
INSERT INTO `ls_chat`.`users_favorites` (id,user_id,message_id) VALUES ('22222222-2222-2222-2222-222222222222','22222222-2222-2222-2222-222222222222','22222222-2222-2222-2222-222222222222');
INSERT INTO `ls_chat`.`users_favorites` (id,user_id,message_id) VALUES ('33333333-3333-3333-3333-333333333333','33333333-3333-3333-3333-333333333333','33333333-3333-3333-3333-333333333333');
INSERT INTO `ls_chat`.`users_favorites` (id,user_id,message_id) VALUES ('44444444-4444-4444-4444-444444444444','44444444-4444-4444-4444-444444444444','44444444-4444-4444-4444-444444444444');
INSERT INTO `ls_chat`.`users_favorites` (id,user_id,message_id) VALUES ('55555555-5555-5555-5555-555555555555','55555555-5555-5555-5555-555555555555','55555555-5555-5555-5555-555555555555');
INSERT INTO `ls_chat`.`users_favorites` (id,user_id,message_id) VALUES ('66666666-6666-6666-6666-666666666666','66666666-6666-6666-6666-666666666666','66666666-6666-6666-6666-666666666666');
-- threads_tags

