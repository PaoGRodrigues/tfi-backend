package services

import (
	"database/sql"

	hosts_domains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	traffic_domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type SQLClient struct {
	db *sql.DB
}

func NewSQLClient(dbConn *sql.DB) *SQLClient {
	return &SQLClient{
		db: dbConn,
	}
}

func (client *SQLClient) AddActiveFlows(flows []traffic_domains.ActiveFlow) error {

	for _, flow := range flows {
		_, err := client.addActiveFlow(flow)
		if err != nil {
			return err
		}
	}
	return nil
}

func (client *SQLClient) addActiveFlow(currentFlow traffic_domains.ActiveFlow) (string, error) {
	_, err := client.db.Exec("INSERT INTO traffic VALUES(?,?,?,?) ON CONFLICT(key) DO UPDATE SET bytes=?;",
		currentFlow.Key, currentFlow.FirstSeen, currentFlow.LastSeen, currentFlow.Bytes, currentFlow.Bytes)
	if err != nil {
		return currentFlow.Key, err
	}

	err = client.insertClient(currentFlow.Client, currentFlow.Key)
	if err != nil {
		return "", err
	}
	err = client.insertServer(currentFlow.Server, currentFlow.Key)
	if err != nil {
		return "", err
	}
	err = client.insertProtocol(currentFlow.Protocol, currentFlow.Key)
	if err != nil {
		return "", err
	}

	return currentFlow.Key, nil
}

func (client *SQLClient) insertClient(currentClient traffic_domains.Client, key string) error {
	_, err := client.db.Exec("INSERT INTO clients VALUES(?,?,?,?);",
		key, currentClient.Name, currentClient.IP, currentClient.Port)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) insertServer(currentServer traffic_domains.Server, key string) error {
	_, err := client.db.Exec("INSERT INTO servers VALUES(?,?,?,?,?,?,?);",
		key, currentServer.Name, currentServer.IP, currentServer.Port, currentServer.IsBroadcastDomain,
		currentServer.IsDHCP, currentServer.Country)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) insertProtocol(currentProto traffic_domains.Protocol, key string) error {
	_, err := client.db.Exec("INSERT INTO protocols VALUES(?,?,?);",
		key, currentProto.L4, currentProto.L7)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) GetServerByAttr(attr string) (traffic_domains.Server, error) {
	server := traffic_domains.Server{}

	rows, err := client.db.Query("SELECT * FROM servers WHERE name LIKE ? LIMIT 1", attr)
	if err != nil {
		return traffic_domains.Server{}, err
	}
	if rows.Next() {
		err = rows.Scan(&server.Key, &server.Name, &server.IP, &server.Port, &server.IsBroadcastDomain, &server.IsDHCP, &server.Country)
		if err != nil {
			return traffic_domains.Server{}, err
		}
	} else {
		rows, err = client.db.Query("SELECT * FROM servers WHERE ip LIKE ? LIMIT 1", attr)
		if err != nil {
			return traffic_domains.Server{}, err
		}
		for rows.Next() {
			err = rows.Scan(&server.Key, &server.Name, &server.IP, &server.Port, &server.IsBroadcastDomain, &server.IsDHCP, &server.Country)
			if err != nil {
				return traffic_domains.Server{}, err
			}
		}
	}

	return server, nil
}

func (client *SQLClient) GetClients() ([]traffic_domains.Client, error) {
	clients := []traffic_domains.Client{}

	rows, err := client.db.Query("SELECT * FROM clients GROUP BY ip")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cli := traffic_domains.Client{}
		err = rows.Scan(&cli.Key, &cli.Name, &cli.IP, &cli.Port)
		if err != nil {
			return nil, err
		}
		clients = append(clients, cli)
	}

	return clients, nil
}

func (client *SQLClient) GetServers() ([]traffic_domains.Server, error) {
	servers := []traffic_domains.Server{}

	rows, err := client.db.Query("SELECT * FROM servers GROUP BY ip")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		srv := traffic_domains.Server{}
		err = rows.Scan(&srv.Key, &srv.Name, &srv.IP, &srv.Port, &srv.IsBroadcastDomain, &srv.IsDHCP, &srv.Country)
		if err != nil {
			return nil, err
		}
		servers = append(servers, srv)
	}

	return servers, nil
}

func (client *SQLClient) GetFlowByKey(key string) (traffic_domains.ActiveFlow, error) {
	flow := traffic_domains.ActiveFlow{}

	rows, err := client.db.Query("SELECT * FROM traffic WHERE key LIKE ? LIMIT 1", key)
	if err != nil {
		return traffic_domains.ActiveFlow{}, err
	}
	for rows.Next() {
		err = rows.Scan(&flow.Key, &flow.FirstSeen, &flow.LastSeen, &flow.Bytes)
		if err != nil {
			return traffic_domains.ActiveFlow{}, err
		}
	}

	return flow, nil
}

func (client *SQLClient) AddHosts(hosts []hosts_domains.Host) error {
	for _, host := range hosts {
		err := client.addHost(host)
		if err != nil {
			return err
		}
	}
	return nil
}

func (client *SQLClient) addHost(host hosts_domains.Host) error {
	_, err := client.db.Exec("INSERT INTO hosts VALUES(?,?,?,?,?,?,?);",
		host.Name, host.ASname, host.PrivateHost, host.IP, host.Mac, host.City, host.Country)
	if err != nil {
		return err
	}
	return nil
}
