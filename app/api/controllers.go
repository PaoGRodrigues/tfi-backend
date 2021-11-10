package api

import (
	"net/http"

	"github.com/PaoGRodrigues/tfi-backend/app/device/handlers"
	"github.com/gin-gonic/gin"
)

type Api struct {
	DeviceHandler *handlers.DeviceHandler
	*gin.Engine
}

func (api *Api) GetDevices(c *gin.Context) {
	devices, err := api.DeviceHandler.GetDevices()
	if err != nil {
		c.JSON(500, gin.H{"data": "error"})
	}
	c.JSON(http.StatusOK, gin.H{"data": devices})
}
