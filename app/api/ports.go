package api

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

type GetLocalhostsUseCase interface {
	GetLocalHosts() ([]host.Host, error)
}
