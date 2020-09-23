package repository

// SQLHandler db処理
type SQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	QueryRow(string, ...interface{}) Row
	Query(string, ...interface{}) (Rows, error)
}

// Result exec result
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Row queryrow result
type Row interface {
	Scan(...interface{}) error
}

// Rows query result
type Rows interface {
	Close() error
	Next() bool
	Scan(...interface{}) error
}
