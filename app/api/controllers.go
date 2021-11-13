package api

import (
	"net/http"

	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	"github.com/gin-gonic/gin"
)

type Api struct {
	DeviceUseCase domains.DeviceUseCase
	*gin.Engine
}

func (api *Api) GetDevices(c *gin.Context) {
	devices, err := api.DeviceUseCase.GetAllDevices()
	if err != nil {
		c.JSON(500, gin.H{"data": "error"})
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": devices})
}
