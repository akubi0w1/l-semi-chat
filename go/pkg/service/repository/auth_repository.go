package repository

import (
	"l-semi-chat/pkg/domain"
)

type authRepository struct {
	SQLHandler SQLHandler
}

// AuthRepository 認証系で使うDBアクセス
type AuthRepository interface {
	FindUserByUserID(userID string) (user domain.User, err error)
}

// NewAuthRepository authRepositoryの作成
func NewAuthRepository(sh SQLHandler) AuthRepository {
	return &authRepository{
		SQLHandler: sh,
	}
}

func (ar *authRepository) FindUserByUserID(userID string) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow("SELECT user_id, password FROM users WHERE user_id=?", userID)
	if err = row.Scan(&user.UserID, &user.Password); err != nil {
		return user, domain.InternalServerError(err)
	}
	return
}
