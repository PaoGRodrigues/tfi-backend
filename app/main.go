package main

import (
	"database/sql"
	"flag"

	alerts_domains "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	alerts_useCases "github.com/PaoGRodrigues/tfi-backend/app/alerts/usecase"
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	hosts_domains "github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	hosts_repository "github.com/PaoGRodrigues/tfi-backend/app/hosts/repository"
	hosts_useCases "github.com/PaoGRodrigues/tfi-backend/app/hosts/usecase"
	services "github.com/PaoGRodrigues/tfi-backend/app/services"
	traffic_domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	traffic_repository "github.com/PaoGRodrigues/tfi-backend/app/traffic/repository"
	traffic_useCases "github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	usecase_hosts "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"
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

	var getLocalhostsUseCase *usecase_hosts.GetLocalhostsUseCase
	var hostBlocker *usecase_hosts.BlockHostUseCase
	var hostsStorage hosts_domains.HostWriter

	var trafficSearcher traffic_domains.TrafficUseCase
	var trafficBytesParser traffic_domains.TrafficBytesParser
	var trafficStorage traffic_domains.TrafficStorage

	var alertsSearcher alerts_domains.AlertUseCase
	var alertSender alerts_domains.AlertsSender
	// ********************************
	// *********** Repository ***********
	var trafficRepo traffic_domains.TrafficRepository
	var hostRepo hosts_domains.HostsRepository
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
			channel = initializedNotifChannel()
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
	hostRepo = initializeHostRepository(database)
	getLocalhostsUseCase, hostsStorage = initializeHostDependencies(tool, hostRepo)

	trafficRepo = initializeTrafficRepository(database)
	trafficSearcher, trafficBytesParser, trafficStorage = initializeTrafficUseCases(tool, trafficRepo, hostsStorage)

	hostBlocker = initializeHostBlockerUseCase(console)

	alertsSearcher = initializeAlertsDependencies(tool)
	alertSender = initializeAlertSender(channel, alertsSearcher)
	// ****************************************

	api := &api.Api{
		Tool: tool,

		GetLocalhostsUseCase: getLocalhostsUseCase,
		BlockHostUseCase:     hostBlocker,
		HostsStorage:         hostsStorage,
		TrafficSearcher:      trafficSearcher,
		TrafficBytesParser:   trafficBytesParser,
		ActiveFlowsStorage:   trafficStorage,
		AlertsSearcher:       alertsSearcher,
		AlertsSender:         alertSender,
		NotifChannel:         channel,
		Engine:               gin.Default(),
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
func initializeHostDependencies(tool services.Tool, hostRepo hosts_domains.HostsRepository) (*usecase_hosts.GetLocalhostsUseCase, hosts_domains.HostWriter) {

	getLocalhostsUseCase := usecase_hosts.NewGetLocalhostsUseCase(tool)
	hostStorage := hosts_useCases.NewHostsStorage(tool, hostRepo)
	return getLocalhostsUseCase, hostStorage
}

func initializeHostBlockerUseCase(console services.Terminal) *usecase_hosts.BlockHostUseCase {
	hostBlocker := usecase_hosts.NewBlockHostUseCase(console)
	return hostBlocker
}

func initializeHostRepository(db services.Database) hosts_domains.HostsRepository {
	hostRepo := hosts_repository.NewHostsRepo(db)
	return hostRepo
}

// *****************************

// *********** Traffic ***********
func initializeTrafficRepository(db services.Database) traffic_domains.TrafficRepository {
	trafficRepo := traffic_repository.NewFlowsRepo(db)
	return trafficRepo
}

func initializeTrafficUseCases(tool services.Tool, repo traffic_domains.TrafficRepository, hostStorage hosts_domains.HostWriter) (traffic_domains.TrafficUseCase,
	traffic_domains.TrafficBytesParser, traffic_domains.TrafficStorage) {

	trafficSearcher := traffic_useCases.NewTrafficSearcher(tool)
	trafficBytesParser := traffic_useCases.NewBytesParser(repo)
	trafficStorage := traffic_useCases.NewFlowsStorage(trafficSearcher, repo, hostStorage)

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
