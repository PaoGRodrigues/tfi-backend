package host

import (
	hostPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
)

type BlockHostUseCase struct {
	blockService hostPorts.HostBlocker
}

func NewBlocker(srv hostPorts.HostBlocker) *BlockHostUseCase {
	return &BlockHostUseCase{
		blockService: srv,
	}
}

func (blocker *BlockHostUseCase) Block(host string) (*string, error) {
	return blocker.blockService.Block(host)
}
