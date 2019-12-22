package repository

type SQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	QueryRow(string, ...interface{}) Row
	Query(string, ...interface{}) (Rows, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
}

type Rows interface {
	Close() error
	Next() bool
	Scan(...interface{}) error
}
