package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	alerts "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	hostPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
	services "github.com/PaoGRodrigues/tfi-backend/app/services"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	hostUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Tool                 services.Tool
	HostUseCase          hostPorts.HostReader
	TrafficSearcher      traffic.TrafficUseCase
	GetLocalhostsUseCase *hostUsecases.GetLocalhostsUseCase
	TrafficBytesParser   traffic.TrafficBytesParser
	ActiveFlowsStorage   traffic.TrafficStorage
	AlertsSearcher       alerts.AlertUseCase
	BlockHostUseCase     *hostUsecases.BlockHostUseCase
	NotifChannel         services.NotificationChannel
	AlertsSender         alerts.AlertsSender
	HostsStorage         *hostUsecases.StoreHostUseCase
	*gin.Engine
}

type AlertsResponse struct {
	Name        string `json:"Nombre"`
	Category    string `json:"Categoría"`
	Time        string `json:"Fecha/hora"`
	Severity    string `json:"Severidad"`
	Source      string `json:"Origen"`
	Destination string `json:"Destino"`
}

type HostsResponse struct {
	Name        string `json:"Nombre,omitempty"`
	PrivateHost bool   `json:"Privado"`
	IP          string `json:"IP"`
	Mac         string `json:"MAC"`
	ASname      string `json:"ASname,omitempty"`
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

func (api *Api) GetActiveFlowsPerDestination(c *gin.Context) {
	activeFlows, err := api.TrafficBytesParser.GetBytesPerDestination()
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
			Time:        alert.Time,
			Category:    alert.Category,
			Severity:    alert.Severity,
			Source:      source,
			Destination: destination,
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
	activeFlows, err := api.TrafficBytesParser.GetBytesPerCountry()
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

func parseHostResponse(hosts []host.Host) []HostsResponse {
	response := []HostsResponse{}

	for _, host := range hosts {
		h := HostsResponse{
			PrivateHost: host.PrivateHost,
			IP:          host.IP,
			Mac:         host.Mac,
		}
		response = append(response, h)
	}
	return response
}

func (api *Api) StoreHosts(c *gin.Context) {
	err := api.HostsStorage.StoreHosts()
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
