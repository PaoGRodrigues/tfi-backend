package main

import (
	"database/sql"

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
	storageClient, err := newDB()
	if err != nil {

	}

	api := &api.Api{
		Tool:                tool,
		HostUseCase:         hostUseCase,
		HostsFilter:         hostsFilter,
		TrafficSearcher:     trafficSearcher,
		ActiveFlowsSearcher: trafficActiveFlowsSearcher,
		Storage:             storageClient,
		Engine:              gin.Default(),
	}

	api.MapURLToPing()
	api.MapGetHostsURL()
	api.MapGetTrafficURL()
	api.MapGetLocalHostsURL()

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

func newTool() *services.Tool {
	tool := services.NewTool("http://192.168.0.16:3000", 2, "XX", "XX")
	return tool
}

func newDB() (*services.DBService, error) {
	SQLiteClient, err := NewSQLite("/file.db")
	if err != nil {
		return nil, err
	}
	databaseConn := trafficRepo.NewSQLClient(SQLiteClient)
	database := services.NewDatabaseService(databaseConn)
	return database, nil
}

func NewSQLite(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	return db, nil
}
