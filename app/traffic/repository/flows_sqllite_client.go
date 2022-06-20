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

func (client *SQLite) InsertActiveFlow(currentFlow domains.ActiveFlow) (int, error) {
	flowKey := currentFlow.Key
	_, err := client.db.Exec("INSERT INTO traffic VALUES(?,?,?,?);",
		currentFlow.Key, currentFlow.FistSeen, currentFlow.LastSeen, currentFlow.Bytes)
	if err != nil {
		return flowKey, err
	}

	err = client.insertClient(currentFlow.Client, flowKey)
	if err != nil {
		return 0, err
	}
	err = client.insertServer(currentFlow.Server, flowKey)
	if err != nil {
		return 0, err
	}
	err = client.insertProtocol(currentFlow.Protocol, flowKey)
	if err != nil {
		return 0, err
	}

	return flowKey, nil
}

func (client *SQLite) insertClient(currentClient domains.Client, key int) error {
	_, err := client.db.Exec("INSERT INTO clients VALUES(?,?,?,?);",
		key, currentClient.Name, currentClient.IP, currentClient.Port)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLite) insertServer(currentServer domains.Server, key int) error {
	_, err := client.db.Exec("INSERT INTO servers VALUES(?,?,?,?,?,?);",
		key, currentServer.Name, currentServer.IP, currentServer.Port, currentServer.IsBroadcastDomain,
		currentServer.IsDHCP)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLite) insertProtocol(currentProto domains.Protocol, key int) error {
	_, err := client.db.Exec("INSERT INTO servers VALUES(?,?,?);",
		key, currentProto.L4, currentProto.L7)
	if err != nil {
		return err
	}
	return nil
}
