package main

import (
	"github.com/PaoGRodrigues/tfi-backend/app/device/gateway"
	"github.com/PaoGRodrigues/tfi-backend/app/device/handlers"
)

func initializeDeviceDependencies() *handlers.DeviceHandler {
	deviceRepo := gateway.NewDeviceRepository()
	deviceGateway := gateway.NewDeviceSearcher(deviceRepo)
	deviceHandler := handlers.NewDeviceHandler(deviceGateway)

	return deviceHandler
}
