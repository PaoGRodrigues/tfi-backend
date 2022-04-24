package main

import (
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	hostDomain "github.com/PaoGRodrigues/tfi-backend/app/host/domains"
	hostRepo "github.com/PaoGRodrigues/tfi-backend/app/host/repository"
	hostUseCase "github.com/PaoGRodrigues/tfi-backend/app/host/usecase"
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
		HostUseCase:    initializeHostDependencies(tool),
		TrafficUseCase: initializeTrafficDependencies(tool),
		Engine:         gin.Default(),
	}

	api.MapURLToPing()
	api.MapGetHostsURL()
	api.MapGetTrafficURL()
	api.MapGetLocalHostsURL()

	api.Run(":8080")
}

func initializeHostDependencies(tool *services_tool.Tool) hostDomain.HostUseCase {
	hostRepo := hostRepo.NewHostClient(tool, "/lua/rest/v2/get/host/custom_data.lua")
	hostSearch := hostUseCase.NewHostSearcher(hostRepo)
	return hostSearch
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
