package repository

import (
	"l-semi-chat/pkg/domain"
)

type accountRepository struct {
	SQLHandler SQLHandler
}

type AccountRepository interface {
	StoreAccount(string, string, string, string, string, string, string) error
	FindAccountByUserID(string) (domain.User, error)
	UpdateAccount(string, string, string, string, string, string, string) error
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

func (ar *accountRepository) UpdateAccount(userID, newUserID, name, mail, image, profile, password string) error {
	// create query
	query := "UPDATE users"
	var values []interface{}
	if newUserID != "" {
		query += " SET user_id=?"
		values = append(values, newUserID)
	}
	if name != "" {
		query += " SET name=?"
		values = append(values, name)
	}
	if mail != "" {
		query += " SET mail=?"
		values = append(values, mail)
	}
	if image != "" {
		query += " SET image=?"
		values = append(values, image)
	}
	if profile != "" {
		query += " SET profile=?"
		values = append(values, profile)
	}
	if password != "" {
		query += " SET password=?"
		values = append(values, password)
	}
	query += " WHERE user_id=?;"
	values = append(values, userID)

	// exec
	_, err := ar.SQLHandler.Execute(query, values...)
	return err

}

func (ar *accountRepository) DeleteAccount(userID string) error {
	_, err := ar.SQLHandler.Execute("DELETE FROM users WHERE user_id=?", userID)
	return err
}
