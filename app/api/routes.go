package api

import (
	"github.com/gin-gonic/gin"
)

func (api *Api) MapURLToPing() {
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (api *Api) MapGetTrafficURL() {
	api.GET("/traffic", api.GetTraffic)
}

func (api *Api) MapGetLocalHostsURL() {
	api.GET("/localhosts", api.GetLocalHosts)
}

func (api *Api) MapGetActiveFlowsPerDestinationURL() {
	api.GET("/activeflowsperdest", api.GetActiveFlowsPerDestination)
}

func (api *Api) MapGetActiveFlowsPerCountryURL() {
	api.GET("/activeflowspercountry", api.GetActiveFlowsPerCountry)
}

func (api *Api) MapStoreActiveFlowsURL() {
	api.POST("/activeflows", api.StoreActiveTraffic)
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
