package services

import (
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
)

func (d *FakeTool) GetAllHosts() ([]domains.Host, error) {
	Hosts := []domains.Host{
		domains.Host{
			Name: "Test1",
			IP:   "13.13.13.13",
		},
		domains.Host{
			Name: "Test2",
			IP:   "14.14.14.14",
		},
		domains.Host{
			Name: "Test3",
			IP:   "15.15.15.15",
		},
		domains.Host{
			Name: "Test4",
			IP:   "16.16.16.16",
		},
	}

	return Hosts, nil
}
