package main

import (
	"database/sql"
	"fmt"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	hostDomain "github.com/PaoGRodrigues/tfi-backend/app/host/domains"
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
	activeFlowsStorage, err := initializeActiveFlowsStorage("/file.db", trafficSearcher)
	if err != nil {
		fmt.Println(err.Error())
	}

	api := &api.Api{
		Tool:                tool,
		HostUseCase:         hostUseCase,
		HostsFilter:         hostsFilter,
		TrafficSearcher:     trafficSearcher,
		ActiveFlowsSearcher: trafficActiveFlowsSearcher,
		ActiveFlowsStorage:  activeFlowsStorage,
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
	hostSearcher := hostUseCase.NewHostSearcher(tool)
	hostsFilter := hostUseCase.NewHostsFilter(hostSearcher)
	return hostSearcher, hostsFilter
}

func initializeTrafficDependencies(tool *services.Tool, hostsFilter hostDomain.HostsFilter) (trafficDomain.TrafficUseCase, trafficDomain.TrafficActiveFlowsSearcher) {
	trafficSearcher := trafficUseCase.NewTrafficSearcher(tool)
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
	tool := services.NewTool("http://192.168.0.13:3000", 2, "xxxx", "xxxx")
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
