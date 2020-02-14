package repository

import (
	"l-semi-chat/pkg/domain"
	"time"
)

type threadRepository struct {
	SQLHandler SQLHandler
}

type threadRepository interface{
	StoreThread(id,name,description,limitUsers,isPublic,createdAt timr.Time,UpdatedAt time.Time)error
	FindThreads()(domain.thread, error)
	FindThreadByID(string)(domain.threads, error)
	UpdateThread(string, string, string, int, int)error
	DeleteThread(string)error
	FindUserByThreadID(string)(domain.Users, error)
	JoinParticipantsByThreadID(string, string, string) error
	FindAccountByID(string)(domain.account error)
	RemoveParticipantsByUserID(string, string)error
	
	StoreUsersThread(string, string, string)error
	FindTagsByUserID(string) (domain.Tags, error)
	FindEvaluationsByUserID(string) (domain.EvaluationScores, error)
}

func NewThreadRepository(sh SQLHandler) ThreadRepository {
	return &threadRepository{
		SQLHandler: sh,
	}
}

func (tr *threadRepository) StoreThread(id,name,description,limitUsers,isPublic,createdAt time.Time,UpdatedAt time.Time)error{
	_, err :=tr.SQLHandler.Execute(
		"INSERT INTO threads(id,name,description,limit_users,is_public,created_at,updated_at) VALUES (?,?,?,?,?,?,?)",
		id,
		name,
		description,
		limitUsers,
		isPublic,
		createdAt,
		updatedAt,
	)
	return domain.InternalServerError(err)
}

func (tr *threadRepository) FindThreads()(threads domain.Threads, err error) {
	rows, err := tr.SQLHandler.Query("
	SELECT threads.id, threads.name, threads.description, threads.limit_users, threads.is_public, threads.created_at, threads.updated_at 
	FROM threads
	INNER JOIN users
	ON threads.user_id = user.id")
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
	row := ar.SQLHandler.QueryRow(
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
	query :="UPDATE threads"
	var values []interface{}
	if name != "" {
		query += "SET name=?"
		values = append(values, name)
	}
	if description != "" {
		query += "SET description=?"
		values = append(values, description)
	}
	if limitUsers != "" {
		query += "SET limit_users=?"
		values = append(values, limitUsers)
	}
	if isPublic != "" {
		query += "SET is_public=?"
		values = append(values, isPublic)
	}
	query += " WHERE id=?;"
	values = append(values, threadID)

	_, err :=tr.SQLHandler.Execute(query, values...)
	return domain.InternalServerError(err)
}

func (tr *threadRepository) FindAccountByID(userID string)(user domain.account, err error) {
	row := ar.SQLHandler.QueryRow("SELECT user_id, name, mail, image, profile FROM users WHERE id=?", userID)
	if err = row.Scan(&user.UserID, &user.Name, &user.Mail, &user.Image, &user.Profile); err != nil {
		return user, domain.InternalServerError(err)
	}
	return user, nil
}

func (tr *threadRepository) DeleteThread(threadID string) error {
	_, err := ar.SQLHandler.Execute("DELETE FROM threads WHERE id=?", threadID)
	return domain.InternalServerError(err)
}

func (tr *threadRepository) FindUserIDByThreadID(threadID string)(users domain.Users, err error) {
	rows, err := tr.SQLHandler.Query(`SELECT user_id, name, mail, image, profile FROM users WHERE id IN (SELECT user_id FROM users_threads WHERE thread_id)`)
	if err != nil {
		logger.Error(err)
		return 
	}
	for rows.Next() {
		var user domain.user
		if err = rows.Scan(&user.UserID, &user.Name, &user.Mail, &user.Image, &user.Profile); err != nil {
			continue
		}
		users = append(user, users)
	}
	return
}

func (tr *threadRepository) JoinParticipantsByThreadID(ID ,threadID, userID string) error {
	isAdmin = 0
	_, err := tr.SQLHandler.Execute(
		"INSERT INTO users_threads(id, user_id, thread_id, is_admin) VALUES (?,?,?,?)",
		ID,
		userID,
		threadID,
		isAdmin,
	)

	return domain.InternalServerError(err)
}

func (tr *threadRepository) RemoveParticipantsByUserID(threadID, userID){
	_, err := ar.SQLHandler.Execute("DELETE FROM users_threads WHERE user_id=? AND thread_id=?",userID, threadID)
	return domain.InternalServerError(err)
}

func (tr *threadRepository) StoreUsersThread(id, threadID, userID string) error {
	var isAdmin = 1
	_, err :=tr.SQLHandler.Execute(
		"INSERT INTO threads(id, user_id, thread_id, is_admin) VALUES (?,?,?,?)",
		id,
		userID,
		threadID,
		isAdmin,
	)
	return domain.InternalServerError(err)
}


func (tr *threadRepository) GetIDByUserID(userID string)(id string, err error) {
	row, err := tr.SQLHandler.Query(`SELECT id FROM users WHERE user_id=? `, userID)
	if err = row.Scan(&id.UserID); err != nil {
		return id.UserID, domain.InternalServerError(err)
	}
	return id.UserID, domain.InternalServerError(err)
}

func (tr *threadRepository) FindTagsByUserID(userID string) (tags domain.Tags, err error) {
	rows, err := ar.SQLHandler.Query(
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
	rows, err := ar.SQLHandler.Query(
		`SELECT evaluation_scores.id, evaluations.item, evaluation_scores.score
		FROM ls_chat.evaluation_scores
		INNER JOIN ls_chat.evaluations
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