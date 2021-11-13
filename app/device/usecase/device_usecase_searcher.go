package usecase

import (
	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
)

type DeviceSearcher struct {
	deviceRepo domains.DeviceRepository
}

func NewDeviceSearcher(repo domains.DeviceRepository) *DeviceSearcher {

	return &DeviceSearcher{
		deviceRepo: repo,
	}
}

func (gw *DeviceSearcher) GetAllDevices() ([]domains.Device, error) {
	res, err := gw.deviceRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}
