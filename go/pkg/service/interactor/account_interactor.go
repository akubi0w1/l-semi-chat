package interactor

import (
	"errors"
	"fmt"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
	"l-semi-chat/pkg/service/repository"
	"time"

	"github.com/google/uuid"
)

type accountInteractor struct {
	AccountRepositoy repository.AccountRepository
	PasswordHandler
}

// AccountInteractor メイン処理
type AccountInteractor interface {
	AddAccount(userID, name, mail, image, profile, password string) (domain.User, error)
	UpdateAccount(userID, newUserID, name, mail, image, profile, password string) (domain.User, error)
	ShowAccount(userID string) (domain.User, error)
	DeleteAccount(userID string) error

	ShowTagsByUserID(userID string) (domain.Tags, error)
	AddAccountTag(userID, tagName, categoryID string) (domain.Tag, error)
	DeleteAccountTag(userID, tagID string) error

	ShowEvaluationScoresByUserID(userID string) (domain.EvaluationScores, error)
}

// NewAccountInteractor accountInteractorを作成
func NewAccountInteractor(ar repository.AccountRepository, ph PasswordHandler) AccountInteractor {
	return &accountInteractor{
		AccountRepositoy: ar,
		PasswordHandler:  ph,
	}
}

func (ai *accountInteractor) AddAccount(userID, name, mail, image, profile, password string) (domain.User, error) {
	var user domain.User
	if userID == "" {
		logger.Warn("create account: userID is empty")
		return user, domain.BadRequest(errors.New("userID is empty"))
	}
	if name == "" {
		logger.Warn("create account: name is empty")
		return user, domain.BadRequest(errors.New("name is empty"))
	}
	if mail == "" {
		logger.Warn("create account: mail is empty")
		return user, domain.BadRequest(errors.New("mail is empty"))
	}
	if password == "" {
		logger.Warn("create account: password is empty")
		return user, domain.BadRequest(errors.New("password is empty"))
	}

	// passwordハッシュ
	hash, err := ai.PasswordHandler.PasswordHash(password)
	if err != nil {
		logger.Error(fmt.Sprintf("create account: %s", err.Error()))
		return user, domain.InternalServerError(err)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		logger.Error(fmt.Sprintf("create account: %s", err.Error()))
		return user, domain.InternalServerError(err)
	}

	// timeの取得
	createdAt := time.Now()

	// dbに突っ込む
	err = ai.AccountRepositoy.StoreAccount(
		id.String(),
		userID,
		name,
		mail,
		image,
		profile,
		hash,
		createdAt,
	)
	if err != nil {
		return user, domain.InternalServerError(err)
	}

	// DONE: 評価の初期化とres用の値の取得
	evaluationScores, err := ai.AccountRepositoy.InitializeEvaluations(user.ID)
	if err != nil {
		return user, err
	}

	user.ID = id.String()
	user.UserID = userID
	user.Name = name
	user.Mail = mail
	user.Image = image
	user.Profile = profile
	user.Evaluations = evaluationScores

	return user, nil
}

func (ai *accountInteractor) ShowAccount(userID string) (domain.User, error) {
	return ai.AccountRepositoy.FindAccountByUserID(userID)
}

func (ai *accountInteractor) UpdateAccount(userID, newUserID, name, mail, image, profile, password string) (user domain.User, err error) {
	// password hash
	var hash string
	if password != "" {
		hash, err = ai.PasswordHandler.PasswordHash(password)
		if err != nil {
			logger.Error(fmt.Sprintf("update account: %s", err.Error()))
			return user, domain.InternalServerError(err)
		}
	}

	err = ai.AccountRepositoy.UpdateAccount(userID, newUserID, name, mail, image, profile, hash)
	if err != nil {
		logger.Error(fmt.Sprintf("update account: %s", err.Error()))
		return
	}

	if newUserID == "" {
		user, err = ai.AccountRepositoy.FindAccountByUserID(userID)
	} else {
		user, err = ai.AccountRepositoy.FindAccountByUserID(newUserID)
	}
	if err != nil {
		logger.Error(fmt.Sprintf("update account: %s", err.Error()))
	}
	return
}

func (ai *accountInteractor) DeleteAccount(userID string) error {
	return ai.AccountRepositoy.DeleteAccount(userID)
}

func (ai *accountInteractor) ShowTagsByUserID(userID string) (domain.Tags, error) {
	return ai.AccountRepositoy.FindTagsByUserID(userID)
}

func (ai *accountInteractor) AddAccountTag(userID, tagName, categoryID string) (tag domain.Tag, err error) {
	if userID == "" {
		logger.Warn("add account tag: userID is empty")
		return tag, domain.BadRequest(errors.New("userID is empty"))
	}
	if tagName == "" {
		logger.Warn("add account tag: tag is empty")
		return tag, domain.BadRequest(errors.New("tag is empty"))
	}
	if categoryID == "" {
		logger.Warn("add account tag: categoryID is empty")
		return tag, domain.BadRequest(errors.New("categoryID is empty"))
	}

	// タグが存在しているかチェック
	tag, err = ai.AccountRepositoy.FindTagByTag(tagName, categoryID)
	var id uuid.UUID
	if err != nil {
		// なければ登録
		id, err = uuid.NewRandom()
		if err != nil {
			logger.Error(fmt.Sprintf("add account tag: %s", err.Error()))
			return tag, domain.InternalServerError(err)
		}

		// TODO: StoreTagとかは、tagHandlerが持っててもいいかも
		err = ai.AccountRepositoy.StoreTag(id.String(), tagName, categoryID)
		if err != nil {
			logger.Error(fmt.Sprintf("add account tag: %s", err.Error()))
			return tag, domain.InternalServerError(err)
		}

		tag, err = ai.AccountRepositoy.FindTagByTag(tagName, categoryID)
		if err != nil {
			logger.Error(fmt.Sprintf("add account tag: %s", err.Error()))
			return tag, domain.InternalServerError(err)
		}
	}

	// store
	id, err = uuid.NewRandom()
	if err != nil {
		logger.Error(fmt.Sprintf("add account tag: %s", err.Error()))
		return tag, domain.InternalServerError(err)
	}
	err = ai.AccountRepositoy.StoreAccountTag(id.String(), userID, tag.ID)
	if err != nil {
		return tag, domain.InternalServerError(err)
	}
	return
}

func (ai *accountInteractor) DeleteAccountTag(userID, tagID string) error {
	return ai.AccountRepositoy.DeleteAccountTag(userID, tagID)
}

func (ai *accountInteractor) ShowEvaluationScoresByUserID(userID string) (domain.EvaluationScores, error) {
	return ai.AccountRepositoy.FindEvaluationsByUserID(userID)
}
