package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"l-semi-chat/pkg/service/repository"
)

type sqlHandler struct {
	DB *sql.DB
}

// NewSQLHandler create new sqlhandler
func NewSQLHandler() repository.SQLHandler {
	conn, err := sql.Open("mysql", "root:password-lsemi@tcp(localhost:3316)/ls_chat")
	if err != nil {
		log.Fatal(err)
	}
	var sh sqlHandler
	sh.DB = conn
	return &sh
}

func (sh *sqlHandler) Execute(query string, args ...interface{}) (repository.Result, error) {
	result, err := sh.DB.Exec(query, args...)
	if err != nil {
		return &sqlResult{}, err
	}
	return &sqlResult{Result: result}, err
}

func (sh *sqlHandler) QueryRow(query string, args ...interface{}) repository.Row {
	row := sh.DB.QueryRow(query, args...)
	return &sqlRow{Row: row}
}

func (sh *sqlHandler) Query(query string, args ...interface{}) (repository.Rows, error) {
	rows, err := sh.DB.Query(query, args...)
	if err != nil {
		return &sql.Rows{}, err
	}
	return &sqlRows{Rows: rows}, err
}

type sqlResult struct {
	Result sql.Result
}

func (r *sqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *sqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type sqlRow struct {
	Row *sql.Row
}

func (r *sqlRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

type sqlRows struct {
	Rows *sql.Rows
}

func (r *sqlRows) Close() error {
	return r.Rows.Close()
}

func (r *sqlRows) Next() bool {
	return r.Rows.Next()
}

func (r *sqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}
