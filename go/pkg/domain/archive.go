package domain

// Archive define archive model
type Archive struct {
	ID       string
	Path     string
	Password string
	IsPublic int
	Thread   Thread
}

// Archives define archives model
type Archives []Archive

// 後で消して
type Thread struct {
	ID          string
	Name        string
	Description string
	LimitUsers  int
	IsPublic    int
	CreatedAt   string
	UpdatedAt   string
	Admin       User
}
