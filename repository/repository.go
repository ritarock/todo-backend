package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const dataSource = "file:data.sqlite?cache=shared&_fk=1"

func NewDB() (*sql.DB, error) {
	return sql.Open("sqlite3", dataSource)
}
