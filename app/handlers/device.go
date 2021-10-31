package handlers

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domains"
	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	DeviceUseCase domains.DeviceUseCase
}

func NewDeviceHandler(rg *gin.RouterGroup, deviceUseCase domains.DeviceUseCase) {

	deviceHandler := &DeviceHandler{
		DeviceUseCase: deviceUseCase,
	}

	rg.GET("devices", deviceHandler.GetDevices)
}

func (dh *DeviceHandler) GetDevices(c *gin.Context) {
	devices, _ := dh.DeviceUseCase.GetAll(c.Request.Context())
	c.JSON(200, gin.H{"data": devices})
}
