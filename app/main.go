package main

import (
	"database/sql"
	"flag"

	alertsDomains "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	alertsUseCases "github.com/PaoGRodrigues/tfi-backend/app/alerts/usecase"
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	hostsDomains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	hostsUseCases "github.com/PaoGRodrigues/tfi-backend/app/hosts/usecase"
	services "github.com/PaoGRodrigues/tfi-backend/app/services"
	trafficDomains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	trafficRepo "github.com/PaoGRodrigues/tfi-backend/app/traffic/repository"
	trafficUseCases "github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	"github.com/coreos/go-iptables/iptables"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var tool services.Tool
	var console services.Terminal
	var channel services.NotificationChannel
	var err error
	scope := flag.String("s", "", "scope")
	flag.Parse()

	var activeFlowsStorage trafficDomains.ActiveFlowsStorage

	hostUseCase, hostsFilter := initializeHostDependencies(tool)
	trafficSearcher := initializeTrafficSearcher(tool)

	if *scope != "prod" {
		tool = services.NewFakeTool()
		console = services.NewFakeConsole()
		//channel = services.NewFakeBot()
		channel = initializedNotifChannel()

		activeFlowsStorage = initializeFakeActiveFlowStorage()

	} else {
		tool = services.NewTool("http://XX:3000", 2, "XX", "XX")
		console, err = initializeConsole()
		if err != nil {
			panic(err.Error())
		}

		activeFlowsStorage, err = initializeActiveFlowsStorage("./file.sqlite", trafficSearcher, hostsFilter)
		if err != nil {
			panic(err.Error())
		}
	}

	trafficActiveFlowsSearcher := initializeTrafficDependencies(activeFlowsStorage)
	alertsSearcher := initializeAlertsDependencies(tool, activeFlowsStorage)
	hostBlocker := initializeHostBlocker(console, activeFlowsStorage)
	channel = initializedNotifChannel()
	alertSender := initializeAlertSender(channel, alertsSearcher)

	api := &api.Api{
		Tool:                tool,
		HostUseCase:         hostUseCase,
		HostsFilter:         hostsFilter,
		TrafficSearcher:     trafficSearcher,
		ActiveFlowsSearcher: trafficActiveFlowsSearcher,
		ActiveFlowsStorage:  activeFlowsStorage,
		AlertsSearcher:      alertsSearcher,
		HostBlocker:         hostBlocker,
		AlertsSender:        alertSender,
		NotifChannel:        channel,
		Engine:              gin.Default(),
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

func initializeHostDependencies(tool services.Tool) (hostsDomains.HostUseCase, hostsDomains.HostsFilter) {
	hostSearcher := hostsUseCases.NewHostSearcher(tool)
	hostsFilter := hostsUseCases.NewHostsFilter(hostSearcher)
	return hostSearcher, hostsFilter
}

func initializeTrafficSearcher(tool services.Tool) trafficDomains.TrafficUseCase {
	trafficSearcher := trafficUseCases.NewTrafficSearcher(tool)
	return trafficSearcher
}

func initializeTrafficDependencies(flowStorage trafficDomains.ActiveFlowsStorage) trafficDomains.TrafficActiveFlowsSearcher {
	trafficActiveFlowsSearcher := trafficUseCases.NewBytesParser(flowStorage)
	return trafficActiveFlowsSearcher
}

func initializeActiveFlowsStorage(file string, trafficSearcher trafficDomains.TrafficUseCase, hostFilter hostsDomains.HostsFilter) (trafficDomains.ActiveFlowsStorage, error) {
	db, err := newDB(file)
	if err != nil {
		return nil, err
	}

	activeFlowsStorage := trafficUseCases.NewFlowsStorage(trafficSearcher, db, hostFilter)
	return activeFlowsStorage, nil
}

func initializeAlertsDependencies(tool services.Tool, trafficStorage trafficDomains.ActiveFlowsStorage) alertsDomains.AlertUseCase {
	alertsSearcher := alertsUseCases.NewAlertSearcher(tool, trafficStorage)
	return alertsSearcher
}

func newDB(file string) (*trafficRepo.SQLClient, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	databaseConn := trafficRepo.NewSQLClient(db)
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

func initializeHostBlocker(console services.Terminal, filter trafficDomains.ActiveFlowsStorage) hostsDomains.HostBlocker {
	hostBlocker := hostsUseCases.NewBlocker(console, filter)
	return hostBlocker
}

func initializeAlertSender(notifier services.NotificationChannel, searcher alertsDomains.AlertUseCase) alertsDomains.AlertsSender {
	alertsSender := alertsUseCases.NewAlertNotifier(notifier, searcher)
	return alertsSender
}

func initializedNotifChannel() services.NotificationChannel {
	telegram := services.NewTelegramInterface()
	return telegram
}

// --------------------
// Fake Initialization
func initializeFakeActiveFlowStorage() *trafficUseCases.FlowsRepository {
	fakeRepo := trafficRepo.NewFakeSQLClient()
	activeFlowsStorage := trafficUseCases.NewFlowsStorage(nil, fakeRepo, nil)
	return activeFlowsStorage
}
