package main

import (
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	deviceDomain "github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	deviceRepo "github.com/PaoGRodrigues/tfi-backend/app/device/repository"
	deviceUseCase "github.com/PaoGRodrigues/tfi-backend/app/device/usecase"
	trafficDomain "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	trafficRepo "github.com/PaoGRodrigues/tfi-backend/app/traffic/repository"
	trafficUseCase "github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	api := &api.Api{
		DeviceUseCase:  initializeDeviceDependencies(),
		TrafficUseCase: initializeTrafficDependencies(),
		Engine:         gin.Default(),
	}

	api.MapURLToPing()
	api.MapGetDevicesURL()
	api.MapGetTrafficURL()

	api.Run(":8080")
}

func initializeDeviceDependencies() deviceDomain.DeviceUseCase {
	deviceRepo := deviceRepo.NewDeviceFakeClient()
	deviceUseCase := deviceUseCase.NewDeviceSearcher(deviceRepo)
	return deviceUseCase
}

func initializeTrafficDependencies() trafficDomain.TrafficUseCase {
	trafficRepo := trafficRepo.NewTrafficFakeClient()
	trafficUseCase := trafficUseCase.NewTrafficSearcher(trafficRepo)
	return trafficUseCase
}
