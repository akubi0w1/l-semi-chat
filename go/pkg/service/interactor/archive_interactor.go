package interactor

import (
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/service/repository"

	"github.com/google/uuid"
)

type archiveInteractor struct {
	ArchiveRepository repository.ArchiveRepository
	PasswordHandler
}

type ArchiveInteractor interface {
	ShowArchive(threadID string) (domain.Archive, error)
	AddArchive(password, threadID string, isPublic int) (domain.Archive, error)
	UpdateArchive(threadID, password string, isPublic int) (domain.Archive, error)
	DeleteArchive(threadID string) error
}

func NewArchiveInteractor(ar repository.ArchiveRepository, ph PasswordHandler) ArchiveInteractor {
	return &archiveInteractor{
		ArchiveRepository: ar,
		PasswordHandler:   ph,
	}
}

func (ai *archiveInteractor) ShowArchive(threadID string) (domain.Archive, error) {
	// TODO: passwordがあればそれのチェック
	return ai.ArchiveRepository.FindArchiveByThreadID(threadID)
}

func (ai *archiveInteractor) AddArchive(password, threadID string, isPublic int) (archive domain.Archive, err error) {
	// TODO: privateならちゃんとpasswordがついてるか

	// password hash
	var hash string
	if password != "" {
		hash, err = ai.PasswordHandler.PasswordHash(password)
		if err != nil {
			return archive, domain.InternalServerError(err)
		}
	}

	// id生成
	id, err := uuid.NewRandom()
	if err != nil {
		return archive, domain.InternalServerError(err)
	}

	// TODO:保留: path作成

	err = ai.ArchiveRepository.StoreArchive(id.String(), "", hash, threadID, isPublic)
	if err != nil {
		return archive, domain.InternalServerError(err)
	}

	// TODO: threadに対応させねば
	archive.ID = id.String()
	archive.Password = hash
	archive.IsPublic = isPublic
	archive.Thread.ID = threadID
	return
}

func (ai *archiveInteractor) UpdateArchive(threadID, password string, isPublic int) (archive domain.Archive, err error) {
	// password hash
	var hash string
	if password != "" {
		hash, err = ai.PasswordHandler.PasswordHash(password)
		if err != nil {
			return archive, domain.InternalServerError(err)
		}
	}

	err = ai.ArchiveRepository.UpdateArchive(threadID, hash, isPublic)
	if err != nil {
		return archive, domain.InternalServerError(err)
	}

	return ai.ArchiveRepository.FindArchiveByThreadID(threadID)
}

func (ai *archiveInteractor) DeleteArchive(threadID string) error {
	return ai.ArchiveRepository.DeleteArchive(threadID)
}
