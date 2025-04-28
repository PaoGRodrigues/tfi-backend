package services

import "database/sql"

type SQLClient struct {
	db *sql.DB
}

func NewSQLClient(dbConn *sql.DB) *SQLClient {
	return &SQLClient{
		db: dbConn,
	}
}
