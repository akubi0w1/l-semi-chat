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

	Evaluations EvaluationScores
}

// Users define users model
type Users []User

// 後で消して
type EvaluationScore struct {
	ID    string
	Item  string
	Score int
}

type EvaluationScores []EvaluationScore
