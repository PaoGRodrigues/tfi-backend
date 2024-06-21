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
	trafficRepository "github.com/PaoGRodrigues/tfi-backend/app/traffic/repository"
	trafficUseCases "github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
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

	var err error
	scope := flag.String("s", "", "scope")
	flag.Parse()

	if *scope != "prod" {
		tool = services.NewFakeTool()
		console = services.NewFakeConsole()
		channel = services.NewFakeBot()
		database = services.NewFakeSQLClient()

	} else {
		tool = services.NewTool("http://XX:3000", 2, "XX", "XX")
		console, err = initializeConsole()
		if err != nil {
			panic(err.Error())
		}
		channel = initializedNotifChannel()
		database, err = newDB("./file.sqlite")
		if err != nil {
			panic(err.Error())
		}
	}

	// *********** Repo & Usecases ***********
	hostUseCase, hostsFilter := initializeHostDependencies(tool)

	trafficRepo := initializeTrafficRepository(database)
	trafficSearcher, trafficBytesParser, trafficStorage := initializeTrafficUseCases(tool, trafficRepo, hostsFilter)

	hostBlocker := initializeHostBlockerUseCase(console, trafficRepo)

	alertsSearcher := initializeAlertsDependencies(tool)
	alertSender := initializeAlertSender(channel, alertsSearcher)
	// ****************************************

	api := &api.Api{
		Tool:               tool,
		HostUseCase:        hostUseCase,
		HostsFilter:        hostsFilter,
		TrafficSearcher:    trafficSearcher,
		TrafficBytesParser: trafficBytesParser,
		ActiveFlowsStorage: trafficStorage,
		AlertsSearcher:     alertsSearcher,
		HostBlocker:        hostBlocker,
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
func initializeHostDependencies(tool services.Tool) (hostsDomains.HostUseCase, hostsDomains.HostsFilter) {
	hostSearcher := hostsUseCases.NewHostSearcher(tool)
	hostsFilter := hostsUseCases.NewHostsFilter(hostSearcher)
	return hostSearcher, hostsFilter
}

func initializeHostBlockerUseCase(console services.Terminal, filter trafficDomains.TrafficRepository) hostsDomains.HostBlocker {
	hostBlocker := hostsUseCases.NewBlocker(console, filter)
	return hostBlocker
}

// *****************************

// *********** Traffic ***********
func initializeTrafficRepository(db services.Database) trafficDomains.TrafficRepository {
	trafficRepo := trafficRepository.NewFlowsRepo(db)
	return trafficRepo
}

func initializeTrafficUseCases(tool services.Tool, repo trafficDomains.TrafficRepository, hostFilter hostsDomains.HostsFilter) (trafficDomains.TrafficUseCase,
	trafficDomains.TrafficBytesParser, trafficDomains.TrafficStorage) {

	trafficSearcher := trafficUseCases.NewTrafficSearcher(tool)
	trafficBytesParser := trafficUseCases.NewBytesParser(repo)
	trafficStorage := trafficUseCases.NewFlowsStorage(trafficSearcher, repo, hostFilter)

	return trafficSearcher, trafficBytesParser, trafficStorage
}

// *******************************

// *********** Alerts ***********
func initializeAlertsDependencies(tool services.Tool) alertsDomains.AlertUseCase {
	alertsSearcher := alertsUseCases.NewAlertSearcher(tool)
	return alertsSearcher
}

func initializeAlertSender(notifier services.NotificationChannel, searcher alertsDomains.AlertUseCase) alertsDomains.AlertsSender {
	alertsSender := alertsUseCases.NewAlertNotifier(notifier, searcher)
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
