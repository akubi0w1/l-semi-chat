package interactor

// PasswordHandler passwordの暗号化、複合
type PasswordHandler interface {
	PasswordHash(pw string) (string, error)
	PasswordVerify(hash, pw string) error
}
