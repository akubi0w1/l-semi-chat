package  domain

type users_thread struct {
	ID        string
	userID    string
	threadID  string
	isPublic  int
	noPermit  int
}