package main

import (
	"github.com/PaoGRodrigues/tfi-backend/app/device/handlers"
	"github.com/gin-gonic/gin"
)

type Api struct {
	deviceHandler *handlers.DeviceHandler
	*gin.Engine
}

func (api *Api) GetDevices(c *gin.Context) {
	api.deviceHandler.GetDevices(c)
}
