package interactor

import (
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/interface/auth"
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

func (ai *authInteractor) Login(userID, password string) error {

	// ユーザの取得
	user, err := ai.AuthRepository.FindUserByUserID(userID)
	if err != nil {
		return err
	}

	// password の比較
	err = auth.PasswordVerify(user.Password, password)

	return domain.Unauthorized(err)
}
