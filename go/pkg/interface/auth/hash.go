package auth

import (
	"l-semi-chat/pkg/service/interactor"

	"golang.org/x/crypto/bcrypt"
)

type passwordHandler struct{}

// NewPasswordHandler パスワードの暗号化、検証を行うためのハンドラ
func NewPasswordHandler() interactor.PasswordHandler {
	return &passwordHandler{}
}

// PasswordHash パスワードのハッシュ
func (ph *passwordHandler) PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// PasswordVerify パスワードの検証
func (ph *passwordHandler) PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
