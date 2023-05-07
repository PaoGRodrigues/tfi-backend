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
	NotifChannel        services.NotificationChannel
	AlertsSender        alerts.AlertsSender
	*gin.Engine
}

type AlertsResponse struct {
	Name        string
	Family      string
	Time        string
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
			Severity:    alert.Severity,
			Source:      source,
			Destination: destination,
			Protocol:    alert.AlertProtocol.Label,
		}
		response = append(response, ar)
	}
	return response
}

func createFlowString(flow alerts.AlertFlow) (string, string) {
	var source strings.Builder
	var destination strings.Builder

	destName := flow.Server.Name

	source.WriteString(flow.Client.IP)
	source.WriteString(":")
	source.WriteString(strconv.Itoa(flow.Client.Port))
	if destName == "" {
		destName = flow.Server.IP
	}
	destination.WriteString(destName)
	destination.WriteString(":")
	destination.WriteString(strconv.Itoa(flow.Server.Port))

	return source.String(), destination.String()
}

type blockHostRequest struct {
	Host string `json:"host" binding:"required"` // Host can be IP or URL
}

func (api *Api) BlockHost(c *gin.Context) {
	host := blockHostRequest{}
	if err := c.BindJSON(&host); err != nil {
		c.JSON(400, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	blockedHost, err := api.HostBlocker.Block(host.Host)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if (hosts.Host{} == blockedHost) {
		c.JSON(400, gin.H{"data": "Host not found"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "POST")
	c.JSON(http.StatusOK, gin.H{"message": "Host has been blocked"})
}

func (api *Api) SendAlertNotification(c *gin.Context) {
	err := api.AlertsSender.SendLastAlertMessages()
	if err != nil {
		c.JSON(500, gin.H{"data": "error getting last alerts"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "POST")
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

type configRequest struct {
	Token    string `json:"token" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func (api *Api) ConfigNotificationChannel(c *gin.Context) {
	config := configRequest{}
	if err := c.BindJSON(&config); err != nil {
		c.JSON(400, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := api.NotifChannel.Configure(config.Token, config.Username)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "POST")
	c.JSON(http.StatusOK, gin.H{"Message": "Channel configured"})
}

func (api *Api) GetActiveFlowsPerCountry(c *gin.Context) {
	activeFlows, err := api.ActiveFlowsSearcher.GetBytesPerCountry()
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
