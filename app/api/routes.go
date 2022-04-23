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

func (api *Api) MapGetHostURL() {
	api.GET("/hosts", api.GetHosts)
}

func (api *Api) MapGetTrafficURL() {
	api.GET("/traffic", api.GetTraffic)
}
