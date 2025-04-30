package api

import (
	"fmt"
	"net/http"

	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	alertUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/alert"
	hostUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"
	notificationChannelUseCases "github.com/PaoGRodrigues/tfi-backend/app/usecase/notificationchannel"
	trafficUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"

	"github.com/gin-gonic/gin"
)

type Api struct {
	TrafficSearcher                     *trafficUsecases.GetTrafficFlowsUseCase
	GetLocalhostsUseCase                *hostUsecases.GetLocalhostsUseCase
	TrafficBytesParser                  traffic.TrafficBytesParser
	StoreTrafficFlowsUseCase            *trafficUsecases.StoreTrafficFlowsUseCase
	GetAlertsUseCase                    *alertUsecases.GetAlertsUseCase
	BlockHostUseCase                    *hostUsecases.BlockHostUseCase
	ConfigureNotificationChannelUseCase *notificationChannelUseCases.ConfigureChannelUseCase
	NotifyAlertsUseCase                 *alertUsecases.NotifyAlertsUseCase
	StoreHostsUseCase                   *hostUsecases.StoreHostUseCase
	*gin.Engine
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
	err := api.StoreTrafficFlowsUseCase.StoreTrafficFlows()
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
