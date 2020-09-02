package repository

import (
	"l-semi-chat/pkg/domain"
	"time"
)

type threadRepository struct {
	SQLHandler SQLHandler
}

type threadRepository interface{
	StoreThread(id,name,description,userID string,limitUsers,isPublic int,createdAt ,UpdatedAt time.Time)error
	FindThreads()(domain.thread, error)
	FindThreadByID(threadID string)(domain.threads, error)
	UpdateThread(threadID, name, description string, limitUsers, isPublic int)error
	DeleteThread(threadID string)error
	FindUsersByThreadID(threadID string)(domain.Users, error)
	JoinParticipantsByThreadID(ID, threadID, userID string) error
	FindAccountByID(userID string)(domain.account error)
	RemoveParticipantsByUserID(threadID, userID string)error
	
	StoreUsersThread(id, threadID, userID string)error
	FindTagsByUserID(userID string) (domain.Tags, error)
	FindEvaluationsByUserID(userID string) (domain.EvaluationScores, error)
	FindUserByThreadID(threadID string)(string, error)
}

func NewThreadRepository(sh SQLHandler) ThreadRepository {
	return &threadRepository{
		SQLHandler: sh,
	}
}

func (tr *threadRepository) StoreThread(id, name, description, userID string, limitUsers, isPublic int, createdAt, UpdatedAt time.Time)error{
	_, err :=tr.SQLHandler.Execute(
		"INSERT INTO threads(id,name,description,limit_users,user_id,is_public,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?)",
		id,
		name,
		description,
		limitUsers,
		userID
		isPublic,
		createdAt,
		updatedAt,
	)
	return domain.InternalServerError(err)
}

func (tr *threadRepository) FindThreads()(threads domain.Threads, err error) {
	rows, err := tr.SQLHandler.Query(`
	SELECT threads.id, threads.name, threads.description, threads.limit_users, threads.is_public, threads.created_at, threads.updated_at, 
	users.id, users.user_id, users.name
	FROM threads
	INNER JOIN users
	ON threads.user_id = users.id`)
	if err != nil {
		logger.Error(fmt.Sprintf("Find thread: %s",err.Error()))
		return threads, domain.InternalServerError(err)
	}

	for rows.Next() {
		var thread domain.Thread
		if err = row.Scan(&thread.ID, &thread.name, &thread.description, &thread.limit_users, &thread.user_id, &thread.is_public, &thread.created_at, &thread.updated_at); err != nil {
			continue
		}
		threads = append(threads, thread)
	}
	
	return threads, domain.InternalServerError(err)
}

func (tr *threadRepository) FindThreadByID(threadID string)(thread domain.Thread, err error) {
	row := tr.SQLHandler.QueryRow(
		`SELECT threads.id, threads.name, threads.description, threads.limit_users, threads.user_id, threads.is_public, threads.created_at, threads.updated_at, users.user_id, users.name, users.image, users.profile
		FROM threads
		INNER JOIN users
		ON users.id = threads.user_id
		WHERE threads.id = ?`,
		threadID,
	)
	if err = row.Scan(&thread.ID, &thread.Name, &thread.Description, &thread.LimitUsers, &thread.Admin.ID, &thread.IsPublic, &thread.CreatedAt, &thread.UpdatedAt, &thread.Admin.UserID, &thread.Admin.Name, &thread.Admin.Image, &thread.Admin.Profile); err != nil {
		return thread, domain.InternalServerError(err)
	}
	return
}

func (tr *threadRepository) UpdateThread(threadID, name, description string, limitUsers, isPublic int)error {
	query :="UPDATE threads SET"
	var values []interface{}
	if name != "" {
		query += "name=?,"
		values = append(values, name)
	}
	if description != "" {
		query += "description=?,"
		values = append(values, description)
	}
	if limitUsers != "" {
		query += "limit_users=?,"
		values = append(values, limitUsers)
	}
	if isPublic != "" {
		query += "is_public=?,"
		values = append(values, isPublic)
	}
	query = strings.TrimSuffix(query, ",")
	query += " WHERE id=?;"
	values = append(values, threadID)

	_, err :=tr.SQLHandler.Execute(query, values...)
	return domain.InternalServerError(err)
}

func (tr *threadRepository) FindAccountByID(userID string)(user domain.User, err error) {
	row := tr.SQLHandler.QueryRow("SELECT user_id, name, mail, image, profile FROM users WHERE id=?", userID)
	if err = row.Scan(&user.UserID, &user.Name, &user.Mail, &user.Image, &user.Profile); err != nil {
		return user, domain.InternalServerError(err)
	}
	return user, nil
}

func (tr *threadRepository) DeleteThread(threadID string) error {
	_, err := tr.SQLHandler.Execute("DELETE FROM threads WHERE id=?", threadID)
	return domain.InternalServerError(err)
}

func (tr *threadRepository) FindUsersByThreadID(threadID string)(users domain.Users, err error) {
	rows, err := tr.SQLHandler.Query(`
		SELECT users.id, users.user_id, users.name, users.image, users.profile
		FROM  users
		INNER JOIN (select user_id from users_threads where thread_id=?) as participants
		ON participants.user_id = users.id`
	)
	if err != nil {
		logger.Error(err)
		return 
	}
	for rows.Next() {
		var user domain.User
		if err = rows.Scan(&user.UserID, &user.Name, &user.Mail, &user.Image, &user.Profile); err != nil {
			continue
		}
		users = append(user, users)
	}
	return
}
JoinParticipantsByThreadID
func (tr *threadRepository) JoinParticipantsByThreadID(ID ,threadID, userID string) error {
	var ut domain.users_thread
	var num := 0
	rows, err := tr.SQLHandler.Query(`SELECT user_id no_permit FROM users_threads WHERE thread_id=?`, threadID)
	if err == nil {
		for rows.Next(){
			if err = rows.Scan(&ut.userID, &ut.noPermit);err != nil {
				continue
			}
			
			if ut.noPermit == 0 {
				num++
			}
			else if ut.userID == userID{
				return domain.Unauthorized(err)
			}
		}
	}
	else {
		var thread domain.thread
		row := tr.SQLHandler.QueryRow(`SELECT limit_users FROM threads WHERE id=?`, threadID)
		if err = row.Scan(&thread.limitUsers); err != nil {
			return domain.InternalServerError(err)
		} 
		if thread.limitUsers<=num {
			return domain.Unauthorized(err)
		}
		_, err = tr.SQLHandler.Execute(
			"INSERT INTO users_threads(id, user_id, thread_id) VALUES (?,?,?)",
			ID,
			userID,
			threadID,
		)
	}
	return domain.InternalServerError(err)
}

func (tr *threadRepository) RemoveParticipantsByUserID(threadID, userID string)error{
	_, err := tr.SQLHandler.Execute("UPDATE users_threads SET no_permit=1 WHERE user_id=? AND thread_id=?",userID, threadID)
	return domain.InternalServerError(err)
}

func (tr *threadRepository) StoreUsersThread(id, threadID, userID string) error {
	_, err :=tr.SQLHandler.Execute(
		"INSERT INTO users_threads(id, user_id, thread_id) VALUES (?,?,?,?)",
		id,
		userID,
		threadID,
	)
	return domain.InternalServerError(err)
}

func (tr *threadRepository) FindTagsByUserID(userID string) (tags domain.Tags, err error) {
	rows, err := tr.SQLHandler.Query(
		`SELECT users_tags.tag_id, tags.tag, tags.category_id, categories.category
		FROM users_tags
		INNER JOIN tags
		ON tags.id = users_tags.tag_id
		INNER JOIN categories
		ON tags.category_id = categories.id
		WHERE user_id = ?`,
		userID,
	)
	if err != nil {
		return tags, domain.InternalServerError(err)
	}
	for rows.Next() {
		var tag domain.Tag
		if err = rows.Scan(&tag.ID, &tag.Tag, &tag.Category.ID, &tag.Category.Category); err != nil {
			continue
		}
		tags = append(tags, tag)
	}
	return
}
func (tr *threadRepository) FindEvaluationsByUserID(userID string) (es domain.EvaluationScores, err error) {
	rows, err := tr.SQLHandler.Query(
		`SELECT evaluation_scores.id, evaluations.item, evaluation_scores.score
		ON evaluations.id = evaluation_scores.evaluation_id
		INNER JOIN evaluations
		ON evaluations.id = evaluation_scores.evaluation_id
		WHERE evaluation_scores.user_id=?`,
		userID,
	)
	if err != nil {
		return
	}

	for rows.Next() {
		var item domain.EvaluationScore
		if err = rows.Scan(&item.ID, &item.Item, &item.Score); err != nil {
			continue
		}
		es = append(es, item)
	}
	return
}

func (tr *threadRepository) FindUserByThreadID(threadID string)(userID string, err error) {
	row, err := tr.SQLHandler.Query(`SELECT id, thread_id, user_id, is_admin, no_permit FROM users_threads WHERE is_admin=1 AND thread_id=?`)
	if err != nil {
		return
	}
	var ut domain.users_thread
	if err = row.Scan(&ut.ID, &ut.userID, &ut.threadID, &ut.isPublic, &ut.noPermit); err != nil {
		return
	}
	userID := us.userID
	return
}