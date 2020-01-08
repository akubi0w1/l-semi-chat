package repository

import (
	"l-semi-chat/pkg/domain"
)

type archiveRepository struct {
	SQLHandler SQLHandler
}

type ArchiveRepository interface {
	FindArchiveByThreadID(threadID string) (domain.Archive, error)
	StoreArchive(archiveID, path, password, threadID string, isPublic int) error
	UpdateArchive(threadID, password string, isPublic int) error
	DeleteArchive(threadID string) error
}

func NewArchiveRepository(sh SQLHandler) ArchiveRepository {
	return &archiveRepository{
		SQLHandler: sh,
	}
}

func (ar *archiveRepository) FindArchiveByThreadID(threadID string) (archive domain.Archive, err error) {
	// TODO: threadに対応させる
	row := ar.SQLHandler.QueryRow("SELECT id, path, password, is_public, thread_id FROM archives WHERE thread_id=?", threadID)
	if err = row.Scan(&archive.ID, &archive.Path, &archive.Password, &archive.IsPublic, &archive.Thread.ID); err != nil {
		return archive, domain.InternalServerError(err)
	}
	return
}

func (ar *archiveRepository) StoreArchive(archiveID, path, password, threadID string, isPublic int) error {
	_, err := ar.SQLHandler.Execute(
		"INSERT INTO archives(id, path, password, is_public, thread_id) VALUES (?,?,?,?,?)",
		archiveID,
		path,
		password,
		isPublic,
		threadID,
	)
	return domain.InternalServerError(err)
}

func (ar *archiveRepository) UpdateArchive(threadID, password string, isPublic int) error {
	query := "UPDATE archives SET"
	var values []interface{}
	if password != "" {
		query += " password=?"
		values = append(values, password)
	}
	query += "is_public=? WHERE thread_id=?"
	values = append(values, isPublic, threadID)

	_, err := ar.SQLHandler.Execute(query, values...)
	return domain.InternalServerError(err)
}

func (ar *archiveRepository) DeleteArchive(threadID string) error {
	_, err := ar.SQLHandler.Execute("DELETE FROM archives WHERE thread_id=?", threadID)
	return domain.InternalServerError(err)
}
