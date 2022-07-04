package services

import (
	"database/sql"
	"strconv"

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

type SQLClient struct {
	db *sql.DB
}

func NewSQLClient(dbConn *sql.DB) *SQLClient {
	return &SQLClient{
		db: dbConn,
	}
}

func getTables() []string {
	return []string{trafficTable, clientsTable, serversTable, protocolsTable}
}

func (client *SQLClient) CreateTables() error {
	tables := getTables()

	for _, table := range tables {
		if _, err := client.db.Exec(table); err != nil {
			return err
		}
	}
	return nil
}

func (client *SQLClient) InsertActiveFlow(currentFlow domains.ActiveFlow) (int, error) {
	flowKey, err := strconv.Atoi(currentFlow.Key)
	if err != nil {
		return 0, err
	}
	_, err = client.db.Exec("INSERT INTO traffic VALUES(?,?,?,?);",
		currentFlow.Key, currentFlow.FistSeen, currentFlow.LastSeen, currentFlow.Bytes)
	if err != nil {
		return flowKey, err
	}

	err = client.InsertClient(currentFlow.Client, flowKey)
	if err != nil {
		return 0, err
	}
	err = client.InsertServer(currentFlow.Server, flowKey)
	if err != nil {
		return 0, err
	}
	err = client.InsertProtocol(currentFlow.Protocol, flowKey)
	if err != nil {
		return 0, err
	}

	return flowKey, nil
}

func (client *SQLClient) InsertClient(currentClient domains.Client, key int) error {
	_, err := client.db.Exec("INSERT INTO clients VALUES(?,?,?,?);",
		key, currentClient.Name, currentClient.IP, currentClient.Port)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) InsertServer(currentServer domains.Server, key int) error {
	_, err := client.db.Exec("INSERT INTO servers VALUES(?,?,?,?,?,?);",
		key, currentServer.Name, currentServer.IP, currentServer.Port, currentServer.IsBroadcastDomain,
		currentServer.IsDHCP)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) InsertProtocol(currentProto domains.Protocol, key int) error {
	_, err := client.db.Exec("INSERT INTO servers VALUES(?,?,?);",
		key, currentProto.L4, currentProto.L7)
	if err != nil {
		return err
	}
	return nil
}
