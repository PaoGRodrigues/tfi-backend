package services

import (
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/coreos/go-iptables/iptables"
)

type Console struct {
	IPTables *iptables.IPTables
}

func NewConsole(ipTablesClient *iptables.IPTables) *Console {
	return &Console{
		IPTables: ipTablesClient,
	}
}

func (c *Console) BlockHost(domains.Host) error {
	return nil
}
