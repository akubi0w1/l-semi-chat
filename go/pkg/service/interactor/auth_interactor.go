package interactor

import (
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/service/repository"
)

type authInteractor struct {
	AuthRepository repository.AuthRepository
	PasswordHandler
}

// AuthInteractor 認証系の処理
type AuthInteractor interface {
	Login(userID, password string) (domain.User, error)
}

// NewAuthInteractor authInteractorの作成
func NewAuthInteractor(ar repository.AuthRepository, ph PasswordHandler) AuthInteractor {
	return &authInteractor{
		AuthRepository:  ar,
		PasswordHandler: ph,
	}
}

func (ai *authInteractor) Login(userID, password string) (domain.User, error) {

	// ユーザの取得
	user, err := ai.AuthRepository.FindUserByUserID(userID)
	if err != nil {
		return user, err
	}

	// password の比較
	err = ai.PasswordHandler.PasswordVerify(user.Password, password)

	return user, domain.Unauthorized(err)
}
