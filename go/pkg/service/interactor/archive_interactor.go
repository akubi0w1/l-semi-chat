package interactor

import (
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/service/repository"

	"github.com/google/uuid"
)

type archiveInteractor struct {
	AccountRepository repository.AccountRepository
	ArchiveRepository repository.ArchiveRepository
	PasswordHandler
}

// ArchiveInteractor archive interactor
type ArchiveInteractor interface {
	ShowArchive(threadID, password string) (domain.Archive, error)
	AddArchive(password, threadID string, isPublic int) (domain.Archive, error)
	UpdateArchive(threadID, password string, isPublic int) (domain.Archive, error)
	DeleteArchive(threadID string) error

	// CheckAdmin threadの管理者であるか確認する
	CheckIsAdmin(threadID, userID string) (bool, error)
}

// NewArchiveInteractor create archive interactor
func NewArchiveInteractor(ar repository.ArchiveRepository, acr repository.AccountRepository, ph PasswordHandler) ArchiveInteractor {
	return &archiveInteractor{
		ArchiveRepository: ar,
		AccountRepository: acr,
		PasswordHandler:   ph,
	}
}

func (ai *archiveInteractor) ShowArchive(threadID, password string) (domain.Archive, error) {
	// passwordがあればそれのチェック
	archive, err := ai.ArchiveRepository.FindArchiveByThreadID(threadID)
	if err != nil {
		return archive, err
	}
	if archive.IsPublic == 0 {
		err = ai.PasswordHandler.PasswordVerify(archive.Password, password)
		if err != nil {
			return domain.Archive{}, err
		}
	}

	user, err := ai.ArchiveRepository.FindUserByID(archive.Thread.Admin.ID)
	if err != nil {
		return archive, err
	}
	archive.Thread.Admin = user
	tags, err := ai.AccountRepository.FindTagsByUserID(user.ID)
	if err != nil {
		return archive, err
	}
	archive.Thread.Admin.Tags = tags
	scores, err := ai.AccountRepository.FindEvaluationsByUserID(user.ID)
	if err != nil {
		return archive, err
	}
	archive.Thread.Admin.Evaluations = scores

	return archive, nil
}

func (ai *archiveInteractor) AddArchive(password, threadID string, isPublic int) (archive domain.Archive, err error) {
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
	path := ""

	err = ai.ArchiveRepository.StoreArchive(id.String(), path, hash, threadID, isPublic)
	if err != nil {
		return archive, domain.InternalServerError(err)
	}

	// TODO: threadに対応させねば
	archive.ID = id.String()
	archive.Path = path
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

func (ai *archiveInteractor) CheckIsAdmin(threadID, userID string) (bool, error) {
	thread, err := ai.ArchiveRepository.FindThreadByThreadID(threadID)
	if err != nil {
		return false, err
	}
	if thread.Admin.UserID != userID {
		return false, nil
	}
	return true, nil
}
