package main

import (
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/device/repository"
	"github.com/PaoGRodrigues/tfi-backend/app/device/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	api := &api.Api{
		DeviceUseCase: initializeDeviceDependencies(),
		Engine:        gin.Default(),
	}

	api.MapURLToPing()
	api.MapGetDevicesURL()

	api.Run(":8080")
}

func initializeDeviceDependencies() domains.DeviceUseCase {
	deviceRepo := repository.NewDeviceFakeClient()
	deviceUseCase := usecase.NewDeviceSearcher(deviceRepo)
	return deviceUseCase
}
