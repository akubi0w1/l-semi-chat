package interactor

import (
	"errors"
	"l-semi-chat/pkg/service/repository"
)

type authInteractor struct {
	AuthRepository repository.AuthRepository
}

type AuthInteractor interface {
	Login(userID, password string) error
}

func NewAuthInteractor(ar repository.AuthRepository) AuthInteractor {
	return &authInteractor{
		AuthRepository: ar,
	}
}

func (ai *authInteractor) Login(userID, password string) (err error) {
	// TODO: bodyのバリデーション

	// ユーザの取得
	user, err := ai.AuthRepository.FindUserByUserID(userID)
	if err != nil {
		return
	}

	// password の比較
	if user.Password != password {
		return errors.New("password valid error")
	}

	return
}
