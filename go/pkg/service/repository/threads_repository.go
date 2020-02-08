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
	FindThread(string)(domain.thread, error)
	UpdateThread
	DeleteThread
}

func NewThreadRepository(sh SQLHandler) ThreadRepository {
	return &threadRepository{
		SQLHandler: sh,
	}
}

func (tr *threadRepository) StoreThread(id,name,description,limitUsers,isPublic,createdAt timr.Time,UpdatedAt time.Time)error{
	_, err :=tr.SQLHandler.Execute(
		"INSERT INTO threads(id,name,description,limit_users,is_public,created_at,updated_at) VALUES (?,?,?,?,?,?,?)"
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
func (tr *threadRepository) FindThread()(thread domain.thread, err error) {
	row := tr.SQLHandler.QueryRow("SELECT id,name,description,limit_users,is_public,created_at,updated_at FROM threads")
	if err = row.Scan(&thread.id, &thread.name, &thread.description, &thread.limit_users, &thread.is_public, &thread.created_at, &thread.updated_at); err != nil {
		return thread, domain.InternalServerError(err)
	}
	return thread, domain.InternalServerError(err)
}