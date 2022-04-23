package main

import (
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	deviceDomain "github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	deviceRepo "github.com/PaoGRodrigues/tfi-backend/app/device/repository"
	deviceUseCase "github.com/PaoGRodrigues/tfi-backend/app/device/usecase"
	services_tool "github.com/PaoGRodrigues/tfi-backend/app/services/tool"
	trafficDomain "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	trafficRepo "github.com/PaoGRodrigues/tfi-backend/app/traffic/repository"
	trafficUseCase "github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	tool := newTool()
	api := &api.Api{
		Tool:           tool,
		DeviceUseCase:  initializeDeviceDependencies(tool),
		TrafficUseCase: initializeTrafficDependencies(tool),
		Engine:         gin.Default(),
	}

	api.MapURLToPing()
	api.MapGetDevicesURL()
	api.MapGetTrafficURL()

	api.Run(":8080")
}

func initializeDeviceDependencies(tool *services_tool.Tool) deviceDomain.DeviceUseCase {
	deviceRepo := deviceRepo.NewDeviceClient(tool, "/lua/rest/v2/get/host/custom_data.lua")
	deviceUseCase := deviceUseCase.NewDeviceSearcher(deviceRepo)
	return deviceUseCase
}

func initializeTrafficDependencies(tool *services_tool.Tool) trafficDomain.TrafficUseCase {
	trafficRepo := trafficRepo.NewActiveFlowClient(tool, "/lua/rest/v2/get/flow/active.lua")
	trafficUseCase := trafficUseCase.NewTrafficSearcher(trafficRepo)
	return trafficUseCase
}

func newTool() *services_tool.Tool {
	tool := services_tool.NewTool("http://192.168.0.16:3000", 2, "XX", "XX")
	return tool
}
