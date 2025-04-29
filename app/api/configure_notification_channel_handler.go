package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	err := api.ConfigureNotificationChannelUseCase.Configure(config.Token, config.Username)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "POST")
	c.JSON(http.StatusOK, gin.H{"Message": "Channel configured"})
}
