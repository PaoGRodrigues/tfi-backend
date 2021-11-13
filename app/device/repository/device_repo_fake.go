package repository

import (
	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
)

type DeviceStorageClient struct {
}

func NewDeviceRepository() *DeviceStorageClient {

	return &DeviceStorageClient{}
}

func (d *DeviceStorageClient) GetAll() ([]domains.Device, error) {
	devices := []domains.Device{
		domains.Device{
			ID:      1,
			Name:    "Test1",
			IP:      "13.13.13.13",
			Details: "details",
		},
		domains.Device{
			ID:      2,
			Name:    "Test2",
			IP:      "14.14.14.14",
			Details: "details",
		},
		domains.Device{
			ID:      3,
			Name:    "Test3",
			IP:      "15.15.15.15",
			Details: "details",
		},
		domains.Device{
			ID:      4,
			Name:    "Test4",
			IP:      "16.16.16.16",
			Details: "details",
		},
	}

	return devices, nil
}
