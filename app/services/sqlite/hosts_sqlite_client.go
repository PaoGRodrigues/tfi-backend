package sqlite

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

func (client *SQLClient) StoreHosts(hosts []host.Host) error {
	for _, host := range hosts {
		err := client.addHost(host)
		if err != nil {
			return err
		}
	}
	return nil
}

func (client *SQLClient) addHost(host host.Host) error {
	_, err := client.db.Exec("INSERT INTO hosts(name,asname,privatehost,ip,mac,city,country) VALUES(?,?,?,?,?,?,?);",
		host.Name, host.ASname, host.PrivateHost, host.IP, host.Mac, host.City, host.Country)
	if err != nil {
		return err
	}
	return nil
}

func (client *SQLClient) GetHostByIp(ip string) (host.Host, error) {
	current := host.Host{}

	rows, err := client.db.Query("SELECT name,asname,privatehost,ip,mac,city,country FROM hosts WHERE ip LIKE ? LIMIT 1", ip)
	if err != nil {
		return host.Host{}, err
	}
	for rows.Next() {
		err = rows.Scan(&current.Name, &current.ASname, &current.PrivateHost, &current.IP, &current.Mac, &current.City, &current.Country)
		if err != nil {
			return host.Host{}, err
		}
	}

	return current, nil
}
