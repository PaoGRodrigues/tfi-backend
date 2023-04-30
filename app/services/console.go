package services

import (
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/coreos/go-iptables/iptables"
)

var chain = "FORWARD"
var table = "filter"

type Console struct {
	IPTables *iptables.IPTables
}

func NewConsole(ipTablesClient *iptables.IPTables) *Console {
	return &Console{
		IPTables: ipTablesClient,
	}
}

func (c *Console) BlockHost(host domains.Host) error {

	exists, err := c.IPTables.ChainExists(table, chain)
	if err != nil {
		return err
	}
	if !exists {
		err := c.IPTables.NewChain(table, chain)
		if err != nil {
			return err
		}
	}

	err = c.IPTables.AppendUnique(table, chain, "-d", host.IP, "-j", "DROP")
	if err != nil {
		return err
	}
	return nil
}
