package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
