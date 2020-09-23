package domain

type Thread struct{
	ID			string
	Name		string
	Description	string
	LimitUsers	string
	UserID		string
	IsPublic	int
	CreatedAt	string
	UpdatedAt 	string
}

type Threads []Thread