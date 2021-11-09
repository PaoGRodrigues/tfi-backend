package gateway

import (
	"context"
	"database/sql"

	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
)

type DeviceStorageClient struct {
	DB *sql.DB
}

func NewDeviceRepository() domains.DeviceRepository {

	return &DeviceStorageClient{}
}

func (d *DeviceStorageClient) GetAll(context.Context) ([]domains.Device, error) {
	return []domains.Device{
		domains.Device{
			ID:      1,
			Name:    "Test",
			IP:      "13.13.13.13",
			Details: "details",
		},
	}, nil
}
