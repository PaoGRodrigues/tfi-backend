package api

import (
	"net/http"

	alertUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/alert"
	hostUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"
	notificationChannelUseCases "github.com/PaoGRodrigues/tfi-backend/app/usecase/notificationchannel"
	trafficUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"
	"github.com/gin-gonic/gin"
)

type Api struct {
	TrafficSearcher                      *trafficUsecases.GetTrafficFlowsUseCase
	GetLocalhostsUseCase                 *hostUsecases.GetLocalhostsUseCase
	GetTrafficFlowsPerDestinationUseCase *trafficUsecases.GetTrafficFlowsPerDestinationUseCase
	GetTrafficFlowsPerCountryUseCase     *trafficUsecases.GetTrafficFlowsPerCountryUseCase
	StoreTrafficFlowsUseCase             *trafficUsecases.StoreTrafficFlowsUseCase
	GetAlertsUseCase                     *alertUsecases.GetAlertsUseCase
	BlockHostUseCase                     *hostUsecases.BlockHostUseCase
	ConfigureNotificationChannelUseCase  *notificationChannelUseCases.ConfigureChannelUseCase
	NotifyAlertsUseCase                  *alertUsecases.NotifyAlertsUseCase
	StoreHostsUseCase                    *hostUsecases.StoreHostUseCase
	*gin.Engine
}

func (api *Api) MapURLToPing() {
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (api *Api) MapSwaggerDocumentation() {
	api.Static("/docs", "./docs")
	// Redirigir a la UI directamente
	api.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/swagger-ui/")
	})

}

func (api *Api) MapGetLocalHostsURL() {
	api.GET("/localhosts", api.GetLocalHosts)
}

func (api *Api) MapGetTrafficURL() {
	api.GET("/traffic", api.GetTraffic)
}

func (api *Api) MapGetActiveFlowsPerDestinationURL() {
	api.GET("/activeflowsperdest", api.GetActiveFlowsPerDestination)
}

func (api *Api) MapGetActiveFlowsPerCountryURL() {
	api.GET("/activeflowspercountry", api.GetActiveFlowsPerCountry)
}

func (api *Api) MapAlertsURL() {
	api.GET("/alerts", api.GetAlerts)
}

func (api *Api) MapBlockHostURL() {
	api.POST("/blockhost", api.BlockHost)
}

func (api *Api) MapNotificationsURL() {
	api.POST("/alertnotification", api.SendAlertNotification)
}

func (api *Api) MapConfigureNotifChannelURL() {
	api.POST("/configurechannel", api.ConfigNotificationChannel)
}

func (api *Api) MapStoreHostsURL() {
	api.POST("/hosts", api.StoreHosts)
}

func (api *Api) MapStoreActiveFlowsURL() {
	api.POST("/activeflows", api.StoreTrafficFlows)
}
