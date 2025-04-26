package repository

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
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

func (hr *HostsRepo) StoreHosts(hosts []host.Host) error {
	err := hr.Database.AddHosts(hosts)
	if err != nil {
		return err
	}
	return nil
}

func (hr *HostsRepo) GetHostByIp(ip string) (host.Host, error) {
	current, err := hr.Database.GetHostByIp(ip)
	if err != nil {
		return host.Host{}, err
	}
	return current, nil
}
