package usecase

import (
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
)

type Blocker struct {
	blockService domains.HostBlockerService
}

func NewBlocker(srv domains.HostBlockerService) *Blocker {
	return &Blocker{
		blockService: srv,
	}
}

func (blocker *Blocker) Block(host string) (*string, error) {
	ahost := host
	err := blocker.blockService.BlockHost(ahost)
	if err != nil {
		return nil, err
	}
	return &ahost, nil
}
