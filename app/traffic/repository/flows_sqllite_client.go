package repository

import (
	"database/sql"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

const trafficTable string = `
	CREATE TABLE IF NOT EXISTS traffic  (
		key INTEGER PRIMARY KEY,
		first_seen INT(10),
		last_seen INT(10),
		bytes BIGINT(20)
	);`

const clientsTable string = `
	CREATE TABLE IF NOT EXISTS clients (
		key INTEGER PRIMARY KEY,
		name TEXT,
		ip VARCHAR(48),
		port UNSIGNED SMALLINT(5),

		FOREIGN KEY (key)
			REFERENCES traffic (key)
			ON DELETE CASCADE
	);`

const serversTable string = `
	CREATE TABLE IF NOT EXISTS clients (
		key INTEGER PRIMARY KEY,
		name TEXT,
		ip VARCHAR(48),
		port UNSIGNED SMALLINT(5),
		is_broadcast_domain BOOLEAN,
		is_dhcp BOOLEAN,

		FOREIGN KEY (key)
			REFERENCES traffic (key)
			ON DELETE CASCADE
	);`

const protocolsTable string = `
	CREATE TABLE IF NOT EXISTS protocols (
		key INTEGER PRIMARY KEY,
		l4 VARCHAR(15),
		l7 VARCHAR(45),

		FOREIGN KEY (key)
			REFERENCES traffic (key)
			ON DELETE CASCADE
	);`

type SQLite struct {
	db *sql.DB
}

func getTables() []string {
	return []string{trafficTable, clientsTable, serversTable, protocolsTable}
}

func NewDatabaseConnection(file string) (*SQLite, error) {

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	tables := getTables()

	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return nil, err
		}
	}

	return &SQLite{
		db: db,
	}, nil
}

func (sqlLite *SQLite) InsertActiveFlow(domains.ActiveFlow) (domains.ActiveFlow, error) {
	return domains.ActiveFlow{}, nil
}
