package main

import (
	"database/sql"
	"flag"

	alerts_domains "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	alerts_useCases "github.com/PaoGRodrigues/tfi-backend/app/alerts/usecase"
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	hosts_domains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	hosts_useCases "github.com/PaoGRodrigues/tfi-backend/app/hosts/usecase"
	services "github.com/PaoGRodrigues/tfi-backend/app/services"
	traffic_domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	traffic_repository "github.com/PaoGRodrigues/tfi-backend/app/traffic/repository"
	traffic_useCases "github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	"github.com/coreos/go-iptables/iptables"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// *********** Services ***********
	var tool services.Tool
	var console services.Terminal
	var channel services.NotificationChannel
	var database services.Database
	// ********************************
	// *********** UseCases ***********
	var hostUseCase hosts_domains.HostUseCase
	var hostsFilter hosts_domains.HostsFilter
	var hostBlocker hosts_domains.HostBlocker

	var trafficSearcher traffic_domains.TrafficUseCase
	var trafficBytesParser traffic_domains.TrafficBytesParser
	var trafficStorage traffic_domains.TrafficStorage

	var alertsSearcher alerts_domains.AlertUseCase
	var alertSender alerts_domains.AlertsSender
	// ********************************
	// *********** Repository ***********
	var trafficRepo traffic_domains.TrafficRepository
	// *******************************

	var err error
	scope := flag.String("s", "", "scope")
	flag.Parse()

	if *scope != "prod" {
		tool = services.NewFakeTool()
		console = services.NewFakeConsole()
		channel = services.NewFakeBot()
		database = services.NewFakeSQLClient()

	} else {
		tool = services.NewTool("http://XXX:3000", 2, "XX", "XX")
		/**
		console, err = initializeConsole()
		if err != nil {
			panic(err.Error())
		}*/
		channel = initializedNotifChannel()
		database, err = newDB("./file.sqlite")
		if err != nil {
			panic(err.Error())
		}
	}

	// *********** Repo & Usecases ***********
	hostUseCase, hostsFilter = initializeHostDependencies(tool)

	trafficRepo = initializeTrafficRepository(database)
	trafficSearcher, trafficBytesParser, trafficStorage = initializeTrafficUseCases(tool, trafficRepo, hostsFilter)

	hostBlocker = initializeHostBlockerUseCase(console, trafficRepo)

	alertsSearcher = initializeAlertsDependencies(tool)
	alertSender = initializeAlertSender(channel, alertsSearcher)
	// ****************************************

	api := &api.Api{
		Tool:               tool,
		HostUseCase:        hostUseCase,
		HostsFilter:        hostsFilter,
		HostBlocker:        hostBlocker,
		TrafficSearcher:    trafficSearcher,
		TrafficBytesParser: trafficBytesParser,
		ActiveFlowsStorage: trafficStorage,
		AlertsSearcher:     alertsSearcher,
		AlertsSender:       alertSender,
		NotifChannel:       channel,
		Engine:             gin.Default(),
	}

	api.MapURLToPing()
	api.MapGetHostsURL()
	api.MapGetTrafficURL()
	api.MapGetLocalHostsURL()
	api.MapGetActiveFlowsPerDestinationURL()
	api.MapStoreActiveFlowsURL()
	api.MapAlertsURL()
	api.MapBlockHostURL()
	api.MapNotificationsURL()
	api.MapConfigureNotifChannelURL()
	api.MapGetActiveFlowsPerCountryURL()

	api.Run(":8080")
}

// *********** Hosts ***********
func initializeHostDependencies(tool services.Tool) (hosts_domains.HostUseCase, hosts_domains.HostsFilter) {
	hostSearcher := hosts_useCases.NewHostSearcher(tool)
	hostsFilter := hosts_useCases.NewHostsFilter(hostSearcher)
	return hostSearcher, hostsFilter
}

func initializeHostBlockerUseCase(console services.Terminal, filter traffic_domains.TrafficRepository) hosts_domains.HostBlocker {
	hostBlocker := hosts_useCases.NewBlocker(console, filter)
	return hostBlocker
}

// *****************************

// *********** Traffic ***********
func initializeTrafficRepository(db services.Database) traffic_domains.TrafficRepository {
	trafficRepo := traffic_repository.NewFlowsRepo(db)
	return trafficRepo
}

func initializeTrafficUseCases(tool services.Tool, repo traffic_domains.TrafficRepository, hostFilter hosts_domains.HostsFilter) (traffic_domains.TrafficUseCase,
	traffic_domains.TrafficBytesParser, traffic_domains.TrafficStorage) {

	trafficSearcher := traffic_useCases.NewTrafficSearcher(tool)
	trafficBytesParser := traffic_useCases.NewBytesParser(repo)
	trafficStorage := traffic_useCases.NewFlowsStorage(trafficSearcher, repo, hostFilter)

	return trafficSearcher, trafficBytesParser, trafficStorage
}

// *******************************

// *********** Alerts ***********
func initializeAlertsDependencies(tool services.Tool) alerts_domains.AlertUseCase {
	alertsSearcher := alerts_useCases.NewAlertSearcher(tool)
	return alertsSearcher
}

func initializeAlertSender(notifier services.NotificationChannel, searcher alerts_domains.AlertUseCase) alerts_domains.AlertsSender {
	alertsSender := alerts_useCases.NewAlertNotifier(notifier, searcher)
	return alertsSender
}

// ******************************

// *********** Services ***********
func initializedNotifChannel() services.NotificationChannel {
	telegram := services.NewTelegramInterface()
	return telegram
}

func newDB(nameFile string) (*services.SQLClient, error) {
	db, err := sql.Open("sqlite3", nameFile)
	if err != nil {
		return nil, err
	}
	databaseConn := services.NewSQLClient(db)
	return databaseConn, nil
}

func initializeConsole() (*services.Console, error) {
	iptables, err := iptables.New()
	if err != nil {
		return nil, err
	}
	console := services.NewConsole(iptables)
	return console, nil
}

// ********************************
