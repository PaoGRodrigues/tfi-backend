package host

import (
	hostPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
)

type BlockHostUseCase struct {
	blockService hostPorts.HostBlocker
}

func NewBlockHostUseCase(blockService hostPorts.HostBlocker) *BlockHostUseCase {
	return &BlockHostUseCase{
		blockService: blockService,
	}
}

func (usecase *BlockHostUseCase) Block(host string) (*string, error) {
	return usecase.blockService.Block(host)
}
