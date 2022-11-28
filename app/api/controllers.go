package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
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
	HostBlocker         hosts.HostBlocker
	*gin.Engine
}

type AlertsResponse struct {
	Name        string
	Family      string
	Time        string
	Score       string
	Severity    string
	Source      string
	Destination string
	Protocol    string
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
	alertsResponse := api.parseAlertsData(alerts)
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": alertsResponse})
}

func (api *Api) parseAlertsData(alerts []domains.Alert) []AlertsResponse {
	response := []AlertsResponse{}
	for _, alert := range alerts {
		source, destination := createFlowString(alert.AlertFlow)
		ar := AlertsResponse{
			Name:        alert.Name,
			Family:      alert.Family,
			Time:        alert.Time.Label,
			Score:       alert.Score,
			Severity:    alert.Severity.Label,
			Source:      source,
			Destination: destination,
			Protocol:    createProtocolString(alert.AlertProtocol),
		}
		response = append(response, ar)
	}
	return response
}

func createFlowString(flow alerts.AlertFlow) (string, string) {
	var source strings.Builder
	var destination strings.Builder

	source.WriteString(flow.Client.Name)
	source.WriteString(":")
	source.WriteString(strconv.Itoa(flow.Client.Port))
	destination.WriteString(flow.Server.Name)
	destination.WriteString(":")
	destination.WriteString(strconv.Itoa(flow.Server.Port))

	return source.String(), destination.String()
}

func createProtocolString(proto alerts.AlertProtocol) string {
	var str strings.Builder

	str.WriteString(proto.Protocol.L4)
	str.WriteString(":")
	str.WriteString(proto.Protocol.L7)

	return str.String()
}

type blockHostRequest struct {
	Ip string `json:"ip" binding:"required"`
}

func (api *Api) BlockHost(c *gin.Context) {
	var host blockHostRequest
	if err := c.ShouldBindJSON(&host); err != nil {
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	_, err := api.HostBlocker.Block(host.Ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "POST")
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
