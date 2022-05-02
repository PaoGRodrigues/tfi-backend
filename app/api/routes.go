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

func (api *Api) MapGetHostsURL() {
	api.GET("/hosts", api.GetHosts)
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
