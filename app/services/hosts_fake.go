package services

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

func (d *FakeTool) GetAllHosts() ([]host.Host, error) {
	Hosts := []host.Host{
		{
			Name:        "Test1",
			IP:          "13.13.13.13",
			PrivateHost: true,
		},
		{
			Name:        "Test2",
			IP:          "14.14.14.14",
			PrivateHost: false,
		},
		{
			Name:        "Test3",
			IP:          "15.15.15.15",
			PrivateHost: true,
		},
		{
			Name:        "Test4",
			IP:          "16.16.16.16",
			PrivateHost: false,
		},
		{
			Name:        "lib.gen.rus",
			IP:          "172.98.98.109",
			PrivateHost: false,
		},
		{
			IP:          "123.123.123.123",
			Name:        "lib.gen.rus",
			PrivateHost: false,
		},
		{
			IP:          "172.98.98.109",
			Name:        "lib.gen.rus",
			PrivateHost: false,
		},
	}

	return Hosts, nil
}
