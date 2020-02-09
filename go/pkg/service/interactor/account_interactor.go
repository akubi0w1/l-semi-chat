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
	AddAccount(string, string, string, string, string, string) (domain.User, error)
	UpdateAccount(string, string, string, string, string, string, string) (domain.User, error)
	ShowAccount(userID string) (domain.User, error)
	DeleteAccount(userID string) error
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

	user.UserID = userID
	user.Name = name
	user.Mail = mail
	user.Image = image
	user.Profile = profile
	return user, nil
}

func (ai *accountInteractor) ShowAccount(userID string) (domain.User, error) {
	return ai.AccountRepositoy.FindAccountByUserID(userID)
}

func (ai *accountInteractor) UpdateAccount(userID, newUserID, name, mail, image, profile, password string) (user domain.User, err error) {
	// password hash
	hash, err := ai.PasswordHandler.PasswordHash(password)
	if err != nil {
		// TODO: ここwarning/400か？
		logger.Error(fmt.Sprintf("update account: %s", err.Error()))
		return user, domain.InternalServerError(err)
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
