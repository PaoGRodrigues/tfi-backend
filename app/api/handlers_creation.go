package api

import (
	"github.com/PaoGRodrigues/tfi-backend/app/device/gateway"
	"github.com/PaoGRodrigues/tfi-backend/app/device/handlers"
)

type Initializer struct{}

func (i *Initializer) InitializeDeviceDependencies() *handlers.DeviceHandler {
	deviceRepo := gateway.NewDeviceRepository()
	deviceGateway := gateway.NewDeviceSearcher(deviceRepo)
	deviceHandler := handlers.NewDeviceHandler(deviceGateway)

	return deviceHandler
}
