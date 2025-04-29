package api

import (
	"fmt"
	"net/http"

	hostPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
	services "github.com/PaoGRodrigues/tfi-backend/app/services"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	alertUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/alert"
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
	GetAlertsUseCase     *alertUsecases.GetAlertsUseCase
	BlockHostUseCase     *hostUsecases.BlockHostUseCase
	NotifChannel         services.NotificationChannel
	NotifyAlertsUseCase  *alertUsecases.NotifyAlertsUseCase
	HostsStorage         *hostUsecases.StoreHostUseCase
	*gin.Engine
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

func (api *Api) SendAlertNotification(c *gin.Context) {
	err := api.NotifyAlertsUseCase.SendAlertMessages()
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
