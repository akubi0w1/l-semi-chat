package repository

import (
	"l-semi-chat/pkg/domain"
	"time"
)

type accountRepository struct {
	SQLHandler SQLHandler
}

// AccountRepository アカウント系で使うDB処理
type AccountRepository interface {
	StoreAccount(id, userID, name, mail, image, profile, password string, createdAt time.Time) error
	FindAccountByUserID(string) (domain.User, error)
	UpdateAccount(string, string, string, string, string, string, string) error
	DeleteAccount(userID string) error

	FindTagsByUserID(userID string) (domain.Tags, error)
	FindTagByTag(tagName, categoryID string) (domain.Tag, error)
	StoreTag(id, tag, categoryID string) error
	StoreAccountTag(id, userID, tagID string) error
	DeleteAccountTag(userID, tagID string) error
}

// NewAccountRepository accountRepositoryの作成
func NewAccountRepository(sh SQLHandler) AccountRepository {
	return &accountRepository{
		SQLHandler: sh,
	}
}

func (ar *accountRepository) StoreAccount(id, userID, name, mail, image, profile, password string, createdAt time.Time) error {
	_, err := ar.SQLHandler.Execute(
		"INSERT INTO users(id, user_id, name, mail, image, profile, password, created_at) VALUES (?,?,?,?,?,?,?,?)",
		id,
		userID,
		name,
		mail,
		image,
		profile,
		password,
		createdAt,
	)
	return domain.InternalServerError(err)
}

func (ar *accountRepository) FindAccountByUserID(userID string) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow("SELECT user_id, name, mail, image, profile FROM users WHERE user_id=?", userID)
	if err = row.Scan(&user.UserID, &user.Name, &user.Mail, &user.Image, &user.Profile); err != nil {
		return user, domain.InternalServerError(err)
	}
	return user, nil
}

func (ar *accountRepository) UpdateAccount(userID, newUserID, name, mail, image, profile, password string) error {
	// create query
	query := "UPDATE users SET"
	var values []interface{}
	if newUserID != "" {
		query += " user_id=?"
		values = append(values, newUserID)
	}
	if name != "" {
		query += " name=?"
		values = append(values, name)
	}
	if mail != "" {
		query += " mail=?"
		values = append(values, mail)
	}
	if image != "" {
		query += " image=?"
		values = append(values, image)
	}
	if profile != "" {
		query += " profile=?"
		values = append(values, profile)
	}
	if password != "" {
		query += " password=?"
		values = append(values, password)
	}
	query += " WHERE user_id=?"
	values = append(values, userID)

	// exec
	_, err := ar.SQLHandler.Execute(query, values...)
	return domain.InternalServerError(err)

}

func (ar *accountRepository) DeleteAccount(userID string) error {
	_, err := ar.SQLHandler.Execute("DELETE FROM users WHERE user_id=?", userID)
	return domain.InternalServerError(err)
}

func (ar *accountRepository) FindTagsByUserID(userID string) (tags domain.Tags, err error) {
	rows, err := ar.SQLHandler.Query(
		`SELECT users_tags.tag_id, tags.tag, tags.category_id, categories.category
		FROM users_tags
		INNER JOIN tags
		ON tags.id = users_tags.tag_id
		INNER JOIN categories
		ON tags.category_id = categories.id
		WHERE user_id = ?`,
		userID,
	)
	if err != nil {
		return tags, domain.InternalServerError(err)
	}
	for rows.Next() {
		var tag domain.Tag
		if err = rows.Scan(&tag.ID, &tag.Tag, &tag.Category.ID, &tag.Category.Category); err != nil {
			continue
		}
		tags = append(tags, tag)
	}
	return
}

func (ar *accountRepository) FindTagByTag(tagName, categoryID string) (tag domain.Tag, err error) {
	row := ar.SQLHandler.QueryRow(
		`SELECT tags.id, tags.tag, tags.category_id, categories.category
		FROM tags
		INNER JOIN categories
		ON tags.category_id = categories.id
		WHERE tags.tag=? and tags.category_id=?`,
		tagName,
		categoryID,
	)
	if err = row.Scan(&tag.ID, &tag.Tag, &tag.Category.ID, &tag.Category.Category); err != nil {
		return tag, domain.InternalServerError(err)
	}
	return
}

func (ar *accountRepository) StoreTag(id, tag, categoryID string) error {
	_, err := ar.SQLHandler.Execute(
		"INSERT INTO tags(id, tag, category_id) VALUES (?,?,?)",
		id,
		tag,
		categoryID,
	)
	return domain.InternalServerError(err)
}

func (ar *accountRepository) StoreAccountTag(id, userID, tagID string) error {
	_, err := ar.SQLHandler.Execute(
		"INSERT INTO users_tags(id, user_id, tag_id) VALUES (?,?,?)",
		id,
		userID,
		tagID,
	)
	return domain.InternalServerError(err)
}

func (ar *accountRepository) DeleteAccountTag(userID, tagID string) error {
	_, err := ar.SQLHandler.Execute("DELETE FROM users_tags WHERE user_id=? and tag_id=?", userID, tagID)
	return domain.InternalServerError(err)
}
