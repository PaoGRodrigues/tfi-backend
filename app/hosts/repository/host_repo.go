package repository

import (
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/services"
)

type HostsRepo struct {
	Database services.Database
}

func NewHostsRepo(database services.Database) *HostsRepo {
	return &HostsRepo{
		Database: database,
	}
}

func (hr *HostsRepo) StoreHosts(hosts []domains.Host) error {
	err := hr.Database.AddHosts(hosts)
	if err != nil {
		return err
	}
	return nil
}

func (hr *HostsRepo) GetHostByIp(ip string) (domains.Host, error) {
	host, err := hr.Database.GetHostByIp(ip)
	if err != nil {
		return domains.Host{}, err
	}
	return host, nil
}
