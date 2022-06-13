package repository

import (
	"database/sql"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type SQLite struct {
	db *sql.DB
}

func NewDatabaseConnection(file string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	return &SQLite{
		db: db,
	}, nil
}

func (db *SQLite) InsertActiveFlow(domains.ActiveFlow) (domains.ActiveFlow, error) {
	return domains.ActiveFlow{}, nil
}
