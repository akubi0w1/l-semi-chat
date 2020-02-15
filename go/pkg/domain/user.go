package domain

// User define user model
type User struct {
	ID          string
	UserID      string
	Name        string
	Mail        string
	Image       string
	Profile     string
	IsAdmin     int
	LoginAt     string
	CreatedAt   string
	Password    string
	Tags        Tags
	Evaluations EvaluationScores
}

// Users define users model
type Users []User

// TODO: 後で消して
type EvaluationScore struct {
	ID    string
	Item  string
	Score int
}

type EvaluationScores []EvaluationScore
