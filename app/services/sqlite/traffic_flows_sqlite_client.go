package sqlite

import (
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
)

func (client *SQLClient) StoreTrafficFlows(flows []traffic.TrafficFlow) error {

	for _, flow := range flows {
		_, err := client.addActiveFlow(flow)
		if err != nil {
			return err
		}
	}
	return nil
}

func (client *SQLClient) addActiveFlow(currentFlow traffic.TrafficFlow) (string, error) {
	_, err := client.db.Exec("INSERT INTO traffic(key,first_seen,last_seen,bytes) VALUES(?,?,?,?) ON CONFLICT(key) DO UPDATE SET bytes=?;",
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

func (client *SQLClient) insertClient(currentClient traffic.Client, key string) error {
	_, err := client.db.Exec("INSERT INTO clients(key,name,ip,port) VALUES(?,?,?,?);",
		key, currentClient.Name, currentClient.IP, currentClient.Port)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) insertServer(currentServer traffic.Server, key string) error {
	_, err := client.db.Exec("INSERT INTO servers(key,name,ip,port,is_broadcast_domain,is_dhcp,country) VALUES(?,?,?,?,?,?,?);",
		key, currentServer.Name, currentServer.IP, currentServer.Port, currentServer.IsBroadcastDomain,
		currentServer.IsDHCP, currentServer.Country)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) insertProtocol(currentProto traffic.Protocol, key string) error {
	_, err := client.db.Exec("INSERT INTO protocols(key,l4,l7) VALUES(?,?,?);",
		key, currentProto.L4, currentProto.L7)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) GetServerByAttr(attr string) (traffic.Server, error) {
	server := traffic.Server{}

	rows, err := client.db.Query("SELECT * FROM servers WHERE name LIKE ? LIMIT 1", attr)
	if err != nil {
		return traffic.Server{}, err
	}
	if rows.Next() {
		err = rows.Scan(&server.Key, &server.Name, &server.IP, &server.Port, &server.IsBroadcastDomain, &server.IsDHCP, &server.Country)
		if err != nil {
			return traffic.Server{}, err
		}
	} else {
		rows, err = client.db.Query("SELECT * FROM servers WHERE ip LIKE ? LIMIT 1", attr)
		if err != nil {
			return traffic.Server{}, err
		}
		for rows.Next() {
			err = rows.Scan(&server.Key, &server.Name, &server.IP, &server.Port, &server.IsBroadcastDomain, &server.IsDHCP, &server.Country)
			if err != nil {
				return traffic.Server{}, err
			}
		}
	}

	return server, nil
}

func (client *SQLClient) GetClients() ([]traffic.Client, error) {
	clients := []traffic.Client{}

	rows, err := client.db.Query("SELECT key,name,ip,port FROM clients GROUP BY key")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cli := traffic.Client{}
		err = rows.Scan(&cli.Key, &cli.Name, &cli.IP, &cli.Port)
		if err != nil {
			return nil, err
		}
		clients = append(clients, cli)
	}

	return clients, nil
}

func (client *SQLClient) GetServers() ([]traffic.Server, error) {
	servers := []traffic.Server{}

	rows, err := client.db.Query("SELECT * FROM servers GROUP BY key")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		srv := traffic.Server{}
		err = rows.Scan(&srv.Key, &srv.Name, &srv.IP, &srv.Port, &srv.IsBroadcastDomain, &srv.IsDHCP, &srv.Country)
		if err != nil {
			return nil, err
		}
		servers = append(servers, srv)
	}

	return servers, nil
}

func (client *SQLClient) GetFlowByKey(key string) (traffic.TrafficFlow, error) {
	flow := traffic.TrafficFlow{}

	rows, err := client.db.Query("SELECT key,first_seen,last_seen,bytes FROM traffic WHERE key LIKE ?", key)
	if err != nil {
		return traffic.TrafficFlow{}, err
	}
	for rows.Next() {
		err = rows.Scan(&flow.Key, &flow.FirstSeen, &flow.LastSeen, &flow.Bytes)
		if err != nil {
			return traffic.TrafficFlow{}, err
		}
	}

	return flow, nil
}
