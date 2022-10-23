package api

import (
	"fmt"
	"net/http"

	alerts "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	hosts "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	services "github.com/PaoGRodrigues/tfi-backend/app/services"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Tool                services.Tool
	HostUseCase         hosts.HostUseCase
	TrafficSearcher     traffic.TrafficUseCase
	HostsFilter         hosts.HostsFilter
	ActiveFlowsSearcher traffic.TrafficActiveFlowsSearcher
	ActiveFlowsStorage  traffic.ActiveFlowsStorage
	AlertsSearcher      alerts.AlertUseCase
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
	traffic, err := api.TrafficSearcher.GetAllActiveTraffic()
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
	hosts, err := api.HostsFilter.GetLocalHosts()
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

func (api *Api) GetActiveFlowsPerDestination(c *gin.Context) {
	activeFlows, err := api.ActiveFlowsSearcher.GetBytesPerDestination()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": activeFlows})
}

func (api *Api) StoreActiveTraffic(c *gin.Context) {
	err := api.ActiveFlowsStorage.StoreFlows()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "POST")
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (api *Api) GetAlerts(c *gin.Context) {
	alerts, err := api.AlertsSearcher.GetAllAlerts()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": alerts})
}
