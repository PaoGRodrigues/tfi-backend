package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

type Blocker struct {
	blockService host.HostBlockerService
}

func NewBlocker(srv host.HostBlockerService) *Blocker {
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
