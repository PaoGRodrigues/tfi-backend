package main

import (
	"database/sql"
	"flag"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	trafficDomains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	alertsPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/alert"
	hostPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
	services "github.com/PaoGRodrigues/tfi-backend/app/services"
	traffic_repository "github.com/PaoGRodrigues/tfi-backend/app/traffic/repository"
	traffic_useCases "github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	alertUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/alert"
	hostUseCases "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"
	notificationChannelUseCases "github.com/PaoGRodrigues/tfi-backend/app/usecase/notificationchannel"
	trafficUseCases "github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"

	"github.com/coreos/go-iptables/iptables"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// *********** Services ***********
	var tool services.Tool
	var console services.Terminal
	var database services.Database
	var channel services.NotificationChannel
	// ********************************
	// *********** UseCases ***********

	var getLocalhostsUseCase *hostUseCases.GetLocalhostsUseCase
	var hostBlocker *hostUseCases.BlockHostUseCase
	var storeHostsUseCase *hostUseCases.StoreHostUseCase

	var trafficSearcher *trafficUseCases.GetTrafficFlowsUseCase
	var trafficBytesParser trafficDomains.TrafficBytesParser
	var trafficStorage trafficDomains.TrafficStorage

	var getAlertsUseCase *alertUsecases.GetAlertsUseCase
	var notifyAlertsUseCase *alertUsecases.NotifyAlertsUseCase

	var configureNotificationChannelUseCase *notificationChannelUseCases.ConfigureChannelUseCase
	// ********************************
	// *********** Repository ***********
	var trafficRepo trafficDomains.TrafficRepository
	// *******************************

	// *********** Flags *************
	var err error
	scope := flag.String("s", "", "scope")
	ip := flag.String("ip", "", "ip")
	port := flag.String("pr", "", "port")
	user := flag.String("u", "", "user")
	pass := flag.String("p", "", "pass")
	db := flag.String("db", "", "db")
	flag.Parse()
	// *******************************

	if *scope != "prod" {
		tool = services.NewFakeTool()
		console = services.NewFakeConsole()
		channel = services.NewFakeBot()
		database = services.NewFakeSQLClient()

	} else {
		if ip != nil || port != nil || user != nil || pass != nil || db != nil {
			tool = services.NewTool("http://"+*ip+":"+*port, *user, *pass)
			err := tool.SetInterfaceID()
			if err != nil {
				panic(err.Error())
			}
			tool.EnableChecks()
			tool.EnableChecks()
			console, err = initializeConsole()
			if err != nil {
				panic(err.Error())
			}
			channel = initializedNotificationChannel()
			database, err = newDB(*db)
			if err != nil {
				panic(err.Error())
			}
		} else {
			if err != nil {
				panic(err.Error())
			}
		}
	}

	// *********** Repo & Usecases ***********
	getLocalhostsUseCase, storeHostsUseCase = initializeHostUseCases(tool, database)

	trafficRepo = initializeTrafficRepository(database)
	trafficSearcher, trafficBytesParser, trafficStorage = initializeTrafficUseCases(tool, trafficRepo, database)

	hostBlocker = initializeHostBlockerUseCases(console)

	configureNotificationChannelUseCase = initializeConfigureNotificationChannelUseCase(channel)

	getAlertsUseCase = initializeGetAlertsUseCases(tool)
	notifyAlertsUseCase = initializeNotifyAlertsUseCases(channel, tool)
	// ****************************************

	api := &api.Api{

		GetLocalhostsUseCase:                getLocalhostsUseCase,
		BlockHostUseCase:                    hostBlocker,
		StoreHostsUseCase:                   storeHostsUseCase,
		TrafficSearcher:                     trafficSearcher,
		TrafficBytesParser:                  trafficBytesParser,
		ActiveFlowsStorage:                  trafficStorage,
		GetAlertsUseCase:                    getAlertsUseCase,
		NotifyAlertsUseCase:                 notifyAlertsUseCase,
		ConfigureNotificationChannelUseCase: configureNotificationChannelUseCase,
		Engine:                              gin.Default(),
	}

	api.MapURLToPing()
	api.MapGetTrafficURL()
	api.MapGetLocalHostsURL()
	api.MapGetActiveFlowsPerDestinationURL()
	api.MapStoreActiveFlowsURL()
	api.MapAlertsURL()
	api.MapBlockHostURL()
	api.MapNotificationsURL()
	api.MapConfigureNotifChannelURL()
	api.MapGetActiveFlowsPerCountryURL()
	api.MapStoreHostsURL()

	api.Run(":8080")
}

// *********** Hosts ***********
func initializeHostUseCases(tool services.Tool, hostDBRepository hostPorts.HostDBRepository) (*hostUseCases.GetLocalhostsUseCase, *hostUseCases.StoreHostUseCase) {

	getLocalhostsUseCase := hostUseCases.NewGetLocalhostsUseCase(tool)
	hostStorage := hostUseCases.NewHostsStorage(tool, hostDBRepository)
	return getLocalhostsUseCase, hostStorage
}

func initializeHostBlockerUseCases(console services.Terminal) *hostUseCases.BlockHostUseCase {
	hostBlocker := hostUseCases.NewBlockHostUseCase(console)
	return hostBlocker
}

// *****************************

// *********** Traffic ***********
func initializeTrafficRepository(db services.Database) trafficDomains.TrafficRepository {
	trafficRepo := traffic_repository.NewFlowsRepo(db)
	return trafficRepo
}

func initializeTrafficUseCases(tool services.Tool, repo trafficDomains.TrafficRepository, hostStorage hostPorts.HostDBRepository) (*trafficUseCases.GetTrafficFlowsUseCase,
	trafficDomains.TrafficBytesParser, trafficDomains.TrafficStorage) {

	getTrafficFlowsUseCase := trafficUseCases.NewTrafficFlowsUseCase(tool)
	trafficBytesParser := traffic_useCases.NewBytesParser(repo)
	trafficStorage := traffic_useCases.NewFlowsStorage(getTrafficFlowsUseCase, repo, hostStorage)

	return getTrafficFlowsUseCase, trafficBytesParser, trafficStorage
}

// *******************************

// *********** Alerts ***********
func initializeGetAlertsUseCases(tool services.Tool) *alertUsecases.GetAlertsUseCase {
	getAlertsUseCase := alertUsecases.NewGetAlertsUseCase(tool)
	return getAlertsUseCase
}

func initializeNotifyAlertsUseCases(notifier alertsPorts.Notifier, tool services.Tool) *alertUsecases.NotifyAlertsUseCase {
	notifyAlertsUseCase := alertUsecases.NewNotifyAlertsUseCase(notifier, tool)
	return notifyAlertsUseCase
}

// ******************************

// *********** Notification Channel ***********
func initializeConfigureNotificationChannelUseCase(channel services.NotificationChannel) *notificationChannelUseCases.ConfigureChannelUseCase {
	configureChannelUseCase := notificationChannelUseCases.NewConfigureChannelUseCase(channel)
	return configureChannelUseCase
}

// ******************************

// *********** Services ***********
func initializedNotificationChannel() services.NotificationChannel {
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
