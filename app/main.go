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

	if *scope != "prod" {
		tool = services.NewFakeTool()
		console = services.NewFakeConsole()
		channel = services.NewFakeBot()
	} else {
		tool = services.NewTool("http://192.168.0.13:3000", 2, "XXX", "XXX")
		console, err = initializeConsole()
		if err != nil {
			panic(err.Error())
		}
		channel = initializedNotifChannel()
	}

	hostUseCase, hostsFilter := initializeHostDependencies(tool)
	trafficSearcher, trafficActiveFlowsSearcher := initializeTrafficDependencies(tool, hostsFilter)
	activeFlowsStorage, err := initializeActiveFlowsStorage("./file.sqlite", trafficSearcher)
	if err != nil {
		panic(err.Error())
	}
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

	api.Run(":8080")
}

func initializeHostDependencies(tool services.Tool) (hostsDomains.HostUseCase, hostsDomains.HostsFilter) {
	hostSearcher := hostsUseCases.NewHostSearcher(tool)
	hostsFilter := hostsUseCases.NewHostsFilter(hostSearcher)
	return hostSearcher, hostsFilter
}

func initializeTrafficDependencies(tool services.Tool, hostsFilter hostsDomains.HostsFilter) (trafficDomains.TrafficUseCase, trafficDomains.TrafficActiveFlowsSearcher) {
	trafficSearcher := trafficUseCases.NewTrafficSearcher(tool)
	trafficActiveFlowsSearcher := trafficUseCases.NewBytesDestinationParser(trafficSearcher, hostsFilter)
	return trafficSearcher, trafficActiveFlowsSearcher
}

func initializeActiveFlowsStorage(file string, trafficSearcher trafficDomains.TrafficUseCase) (trafficDomains.ActiveFlowsStorage, error) {
	db, err := newDB(file)
	if err != nil {
		return nil, err
	}

	activeFlowsStorage := trafficUseCases.NewFlowsStorage(trafficSearcher, db)
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
