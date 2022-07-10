package main

import (
	"database/sql"
	"fmt"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	hostDomain "github.com/PaoGRodrigues/tfi-backend/app/host/domains"
	hostRepo "github.com/PaoGRodrigues/tfi-backend/app/host/repository"
	hostUseCase "github.com/PaoGRodrigues/tfi-backend/app/host/usecase"
	services "github.com/PaoGRodrigues/tfi-backend/app/services"
	trafficDomain "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	trafficRepo "github.com/PaoGRodrigues/tfi-backend/app/traffic/repository"
	trafficUseCase "github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	tool := newTool()
	hostUseCase, hostsFilter := initializeHostDependencies(tool)
	trafficSearcher, trafficActiveFlowsSearcher := initializeTrafficDependencies(tool, hostsFilter)
	storageClient, err := initializeActiveFlowsStorage("/file.db", trafficSearcher)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	api := &api.Api{
		Tool:                tool,
		HostUseCase:         hostUseCase,
		HostsFilter:         hostsFilter,
		TrafficSearcher:     trafficSearcher,
		ActiveFlowsSearcher: trafficActiveFlowsSearcher,
		ActiveFlowsStorage:  storageClient,
		Engine:              gin.Default(),
	}

	api.MapURLToPing()
	api.MapGetHostsURL()
	api.MapGetTrafficURL()
	api.MapGetLocalHostsURL()
	api.MapGetActiveFlowsPerDestinationURL()
	api.MapStoreActiveFlows()

	api.Run(":8080")
}

func initializeHostDependencies(tool *services.Tool) (hostDomain.HostUseCase, hostDomain.HostsFilter) {
	hostRepo := hostRepo.NewHostClient(tool, "/lua/rest/v2/get/host/custom_data.lua")
	hostSearcher := hostUseCase.NewHostSearcher(hostRepo)
	hostsFilter := hostUseCase.NewHostsFilter(hostSearcher)
	return hostSearcher, hostsFilter
}

func initializeTrafficDependencies(tool *services.Tool, hostsFilter hostDomain.HostsFilter) (trafficDomain.TrafficUseCase, trafficDomain.TrafficActiveFlowsSearcher) {
	trafficRepo := trafficRepo.NewActiveFlowClient(tool, "/lua/rest/v2/get/flow/active.lua")
	trafficSearcher := trafficUseCase.NewTrafficSearcher(trafficRepo)
	trafficActiveFlowsSearcher := trafficUseCase.NewBytesDestinationParser(trafficSearcher, hostsFilter)
	return trafficSearcher, trafficActiveFlowsSearcher
}

func initializeActiveFlowsStorage(file string, trafficSearcher trafficDomain.TrafficUseCase) (trafficDomain.ActiveFlowsStorage, error) {
	db, err := newDB(file)
	if err != nil {
		return nil, err
	}

	activeFlowsStorage := trafficUseCase.NewFlowsStorage(trafficSearcher, db)
	return activeFlowsStorage, nil
}

func newTool() *services.Tool {
	tool := services.NewTool("http://192.168.0.13:3000", 2, "admin", "admin")
	return tool
}

func newDB(file string) (*trafficRepo.SQLClient, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	databaseConn := trafficRepo.NewSQLClient(db)
	return databaseConn, nil
}
