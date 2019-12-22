package repository

import "l-semi-chat/pkg/domain"

type accountRepository struct {
	SQLHandler SQLHandler
}

type AccountRepository interface {
	StoreAccount(string, string, string, string, string, string, string) error
	FindAccountByUserID(string) (domain.User, error)
	DeleteAccount(userID string) error
}

func NewAccountRepository(sh SQLHandler) AccountRepository {
	return &accountRepository{
		SQLHandler: sh,
	}
}

func (ar *accountRepository) StoreAccount(id, userID, name, mail, image, profile, password string) error {
	_, err := ar.SQLHandler.Execute(
		"INSERT INTO users(id, user_id, name, mail, image, profile, password) VALUES (?,?,?,?,?,?,?)",
		id,
		userID,
		name,
		mail,
		image,
		profile,
		password,
	)
	return err
}

func (ar *accountRepository) FindAccountByUserID(userID string) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow("SELECT user_id, name, mail, image, profile FROM users WHERE user_id=?", userID)
	if err = row.Scan(&user.UserID, &user.Name, &user.Mail, &user.Image, &user.Profile); err != nil {
		return
	}
	return
}

func (ar *accountRepository) DeleteAccount(userID string) error {
	_, err := ar.SQLHandler.Execute("DELETE FROM users WHERE user_id=?", userID)
	return err
}
