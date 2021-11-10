package handlers

import (
	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
)

type DeviceHandler struct {
	DeviceGateway domains.DeviceGateway
}

func NewDeviceHandler(deviceGateway domains.DeviceGateway) *DeviceHandler {

	return &DeviceHandler{
		DeviceGateway: deviceGateway,
	}
}

func (dh *DeviceHandler) GetDevices() ([]domains.Device, error) {
	devices, err := dh.DeviceGateway.GetAll()

	if err != nil {
		return nil, err
	}

	return devices, nil
}
