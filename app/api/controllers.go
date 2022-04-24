package api

import (
	"fmt"
	"net/http"

	host "github.com/PaoGRodrigues/tfi-backend/app/host/domains"
	services_tool "github.com/PaoGRodrigues/tfi-backend/app/services/tool"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Tool            *services_tool.Tool
	HostUseCase     host.HostUseCase
	TrafficUseCase  traffic.TrafficUseCase
	LocalHostFilter host.LocalHostFilter
	*gin.Engine
}

func (api *Api) GetHosts(c *gin.Context) {
	hosts, err := api.HostUseCase.GetAllHosts()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": hosts})
	return
}

func (api *Api) GetTraffic(c *gin.Context) {
	traffic, err := api.TrafficUseCase.GetAllActiveTraffic()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": traffic})
}

func (api *Api) GetLocalHosts(c *gin.Context) {
	hosts, err := api.LocalHostFilter.GetLocalHosts()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": hosts})
}
