package repository

type accountRepository struct {
	SQLHandler SQLHandler
}

type AccountRepository interface {
	StoreAccount(string, string, string, string, string, string, string) error
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
