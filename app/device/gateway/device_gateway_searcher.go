package gateway

import (
	"context"

	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
)

type DeviceSearcher struct {
	deviceRepo domains.DeviceRepository
}

func NewDeviceSearcher(repo domains.DeviceRepository) domains.DeviceGateway {

	return &DeviceSearcher{
		deviceRepo: repo,
	}
}

func (gw *DeviceSearcher) GetAll(c context.Context) ([]domains.Device, error) {
	res, err := gw.deviceRepo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}
