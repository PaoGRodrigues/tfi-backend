package handlers

import (
	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	DeviceGateway domains.DeviceGateway
}

func NewDeviceHandler(deviceGateway domains.DeviceGateway) *DeviceHandler {

	return &DeviceHandler{
		DeviceGateway: deviceGateway,
	}
}

func (dh *DeviceHandler) GetDevices(c *gin.Context) {
	devices, _ := dh.DeviceGateway.GetAll(c.Request.Context())
	c.JSON(200, gin.H{"data": devices})
}
