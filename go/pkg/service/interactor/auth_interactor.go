package interactor

import (
	"errors"
	"l-semi-chat/pkg/service/repository"

	"l-semi-chat/pkg/interface/auth"
)

type authInteractor struct {
	AuthRepository repository.AuthRepository
}

type AuthInteractor interface {
	Login(userID, password string) (string, error)
}

func NewAuthInteractor(ar repository.AuthRepository) AuthInteractor {
	return &authInteractor{
		AuthRepository: ar,
	}
}

func (ai *authInteractor) Login(userID, password string) (tokenString string, err error) {
	// TODO: bodyのバリデーション

	// ユーザの取得
	user, err := ai.AuthRepository.FindUserByUserID(userID)
	if err != nil {
		return
	}

	// password の比較
	if user.Password != password {
		return "", errors.New("password valid error")
	}

	// jwt発行
	// TODO: 依存してるぅ
	tokenString, err = auth.CreateToken(userID)
	return
}
