package main

import (
	"database/sql"
	"flag"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	alertsPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/alert"
	hostPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
	trafficPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/traffic"
	consoleService "github.com/PaoGRodrigues/tfi-backend/app/services/console"
	ntopngService "github.com/PaoGRodrigues/tfi-backend/app/services/ntopng"
	sqliteService "github.com/PaoGRodrigues/tfi-backend/app/services/sqlite"
	telegramService "github.com/PaoGRodrigues/tfi-backend/app/services/telegram"

	services "github.com/PaoGRodrigues/tfi-backend/app/services"
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

	var getTrafficFlowsUseCase *trafficUseCases.GetTrafficFlowsUseCase
	var getTrafficFlowsPerDestinationUseCase *trafficUseCases.GetTrafficFlowsPerDestinationUseCase
	var getTrafficFlowsPerCountryUseCase *trafficUseCases.GetTrafficFlowsPerCountryUseCase

	var storeTrafficFlowsUseCase *trafficUseCases.StoreTrafficFlowsUseCase

	var getAlertsUseCase *alertUsecases.GetAlertsUseCase
	var notifyAlertsUseCase *alertUsecases.NotifyAlertsUseCase

	var configureNotificationChannelUseCase *notificationChannelUseCases.ConfigureChannelUseCase
	// ********************************

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
		tool = ntopngService.NewFakeTool()
		console = consoleService.NewFakeConsole()
		channel = telegramService.NewFakeBot()
		database = sqliteService.NewFakeSQLClient()

	} else {
		if ip != nil || port != nil || user != nil || pass != nil || db != nil {
			tool = ntopngService.NewTool("http://"+*ip+":"+*port, *user, *pass)
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

	getTrafficFlowsUseCase, getTrafficFlowsPerDestinationUseCase, storeTrafficFlowsUseCase = initializeTrafficUseCases(tool, database, database)

	hostBlocker = initializeHostBlockerUseCases(console)

	configureNotificationChannelUseCase = initializeConfigureNotificationChannelUseCase(channel)

	getAlertsUseCase = initializeGetAlertsUseCases(tool)
	notifyAlertsUseCase = initializeNotifyAlertsUseCases(channel, tool)
	// ****************************************

	api := &api.Api{

		GetLocalhostsUseCase:                 getLocalhostsUseCase,
		BlockHostUseCase:                     hostBlocker,
		StoreHostsUseCase:                    storeHostsUseCase,
		TrafficSearcher:                      getTrafficFlowsUseCase,
		GetTrafficFlowsPerDestinationUseCase: getTrafficFlowsPerDestinationUseCase,
		GetTrafficFlowsPerCountryUseCase:     getTrafficFlowsPerCountryUseCase,
		StoreTrafficFlowsUseCase:             storeTrafficFlowsUseCase,
		GetAlertsUseCase:                     getAlertsUseCase,
		NotifyAlertsUseCase:                  notifyAlertsUseCase,
		ConfigureNotificationChannelUseCase:  configureNotificationChannelUseCase,
		Engine:                               gin.Default(),
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
func initializeTrafficUseCases(tool services.Tool, repo trafficPorts.TrafficDBRepository, hostStorage hostPorts.HostDBRepository) (*trafficUseCases.GetTrafficFlowsUseCase,
	*trafficUseCases.GetTrafficFlowsPerDestinationUseCase, *trafficUseCases.StoreTrafficFlowsUseCase) {

	getTrafficFlowsUseCase := trafficUseCases.NewTrafficFlowsUseCase(tool)
	storeTrafficFlowsUseCase := trafficUseCases.NewStoreTrafficFlowsUseCase(getTrafficFlowsUseCase, repo, hostStorage)
	getTrafficFlowsPerDestinationUseCase := trafficUseCases.NewGetTrafficFlowsPerDestinationUseCase(repo)

	return getTrafficFlowsUseCase, getTrafficFlowsPerDestinationUseCase, storeTrafficFlowsUseCase
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
	telegram := telegramService.NewTelegramInterface()
	return telegram
}

func newDB(nameFile string) (services.Database, error) {
	db, err := sql.Open("sqlite3", nameFile)
	if err != nil {
		return nil, err
	}
	databaseConn := sqliteService.NewSQLClient(db)
	return databaseConn, nil
}

func initializeConsole() (services.Terminal, error) {
	iptables, err := iptables.New()
	if err != nil {
		return nil, err
	}
	console := consoleService.NewConsole(iptables)
	return console, nil
}

// ********************************
