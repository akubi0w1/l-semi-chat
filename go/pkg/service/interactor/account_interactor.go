package interactor

import (
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/interface/auth"
	"l-semi-chat/pkg/service/repository"

	"github.com/google/uuid"
)

type accountInteractor struct {
	AccountRepositoy repository.AccountRepository
}

type AccountInteractor interface {
	AddAccount(string, string, string, string, string, string) (domain.User, error)
	UpdateAccount(string, string, string, string, string, string, string) (domain.User, error)
	ShowAccount(userID string) (domain.User, error)
	DeleteAccount(userID string) error
}

func NewAccountInteractor(ar repository.AccountRepository) AccountInteractor {
	return &accountInteractor{
		AccountRepositoy: ar,
	}
}

func (ai *accountInteractor) AddAccount(userID, name, mail, image, profile, password string) (domain.User, error) {
	// TODO: バリデーションチェック

	// passwordハッシュ
	hash, err := auth.PasswordHash(password)
	if err != nil {
		return domain.User{}, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return domain.User{}, err
	}

	// TODO: timeの取得

	// dbに突っ込む
	err = ai.AccountRepositoy.StoreAccount(
		id.String(),
		userID,
		name,
		mail,
		image,
		profile,
		hash,
	)

	return domain.User{
		UserID:  userID,
		Name:    name,
		Mail:    mail,
		Image:   image,
		Profile: profile,
	}, err
}

func (ai *accountInteractor) ShowAccount(userID string) (domain.User, error) {
	return ai.AccountRepositoy.FindAccountByUserID(userID)
}

func (ai *accountInteractor) UpdateAccount(userID, newUserID, name, mail, image, profile, password string) (user domain.User, err error) {
	// TODO: password hash

	err = ai.AccountRepositoy.UpdateAccount(userID, newUserID, name, mail, image, profile, password)
	if err != nil {
		return
	}
	if newUserID == "" {
		user, err = ai.AccountRepositoy.FindAccountByUserID(userID)
	} else {
		user, err = ai.AccountRepositoy.FindAccountByUserID(newUserID)
	}
	return
}

func (ai *accountInteractor) DeleteAccount(userID string) error {
	return ai.AccountRepositoy.DeleteAccount(userID)
}
