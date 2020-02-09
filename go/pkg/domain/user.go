package domain

// User define user model
type User struct {
	ID        string
	UserID    string
	Name      string
	Mail      string
	Image     string
	Profile   string
	IsAdmin   int
	LoginAt   string
	CreatedAt string
	Password  string
}

// Users define users model
type Users []User
