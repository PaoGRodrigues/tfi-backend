package services

import "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"

type Console struct {
}

func NewConsole() *Console {
	return &Console{}
}

func (c *Console) BlockHost(domains.Host) error {
	return nil
}
