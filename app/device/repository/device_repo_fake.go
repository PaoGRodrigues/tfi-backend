package repository

import (
	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
)

type DeviceFakeClient struct {
}

func NewDeviceFakeClient() *DeviceFakeClient {

	return &DeviceFakeClient{}
}

func (d *DeviceFakeClient) GetAll() ([]domains.Device, error) {
	devices := []domains.Device{
		domains.Device{
			Name:     "Test1",
			IP:       "13.13.13.13",
			OsDetail: "details",
		},
		domains.Device{
			Name:     "Test2",
			IP:       "14.14.14.14",
			OsDetail: "details",
		},
		domains.Device{
			Name:     "Test3",
			IP:       "15.15.15.15",
			OsDetail: "details",
		},
		domains.Device{
			Name:     "Test4",
			IP:       "16.16.16.16",
			OsDetail: "details",
		},
	}

	return devices, nil
}
