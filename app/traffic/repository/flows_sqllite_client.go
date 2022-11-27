package repository

import (
	"database/sql"
	"strconv"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type SQLClient struct {
	db *sql.DB
}

func NewSQLClient(dbConn *sql.DB) *SQLClient {
	return &SQLClient{
		db: dbConn,
	}
}

func (client *SQLClient) AddActiveFlows(flows []domains.ActiveFlow) error {

	for _, flow := range flows {
		_, err := client.addActiveFlow(flow)
		if err != nil {
			return err
		}
	}
	return nil
}

func (client *SQLClient) addActiveFlow(currentFlow domains.ActiveFlow) (int, error) {
	flowKey, err := strconv.Atoi(currentFlow.Key)
	if err != nil {
		return 0, err
	}
	_, err = client.db.Exec("INSERT INTO traffic VALUES(?,?,?,?) ON CONFLICT(key) DO UPDATE SET bytes=?;",
		currentFlow.Key, currentFlow.FirstSeen, currentFlow.LastSeen, currentFlow.Bytes, currentFlow.Bytes)
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

func (client *SQLClient) insertClient(currentClient domains.Client, key int) error {
	_, err := client.db.Exec("INSERT INTO clients VALUES(?,?,?,?);",
		key, currentClient.Name, currentClient.IP, currentClient.Port)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) insertServer(currentServer domains.Server, key int) error {
	_, err := client.db.Exec("INSERT INTO servers VALUES(?,?,?,?,?,?);",
		key, currentServer.Name, currentServer.IP, currentServer.Port, currentServer.IsBroadcastDomain,
		currentServer.IsDHCP)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) insertProtocol(currentProto domains.Protocol, key int) error {
	_, err := client.db.Exec("INSERT INTO protocols VALUES(?,?,?);",
		key, currentProto.L4, currentProto.L7)
	if err != nil {
		return err
	}
	return nil
}
