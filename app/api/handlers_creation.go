package api

import (
	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/device/repository"
	"github.com/PaoGRodrigues/tfi-backend/app/device/usecase"
)

type Initializer struct{}

func (i *Initializer) InitializeDeviceDependencies() domains.DeviceUseCase {
	deviceRepo := repository.NewDeviceRepository()
	deviceUseCase := usecase.NewDeviceSearcher(deviceRepo)

	return deviceUseCase
}
