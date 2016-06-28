package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Messageid int
	Name      string
	Body      string
}

type Repo struct {
	DB *DB
}

func Open(dsn string) (*Repo, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Repo{DB: db}, nil
}
