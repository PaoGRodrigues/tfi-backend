package gateway

import (
	"context"

	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
)

type deviceRepository struct {
}

func NewDeviceRepository() domains.DeviceRepository {

	return &deviceRepository{}
}

func (d *deviceRepository) GetAll(context.Context) ([]domains.Device, error) {
	return []domains.Device{}, nil
}
