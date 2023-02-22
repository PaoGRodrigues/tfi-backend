package services

import (
	"fmt"

	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
)

type FakeConsole struct {
}

func NewFakeConsole() *FakeConsole {
	return &FakeConsole{}
}

func (fc *FakeConsole) BlockHost(host domains.Host) error {
	fmt.Printf("Blocking... %s : %s", host.IP, host.Name)
	return nil
}
