package useCases

import (
	"context"

	"github.com/PaoGRodrigues/tfi-backend/app/domains"
)

type deviceUseCase struct {
	deviceRepo domains.DeviceRepository
}

func NewDeviceUseCase(repo domains.DeviceRepository) domains.DeviceUseCase {

	return &deviceUseCase{
		deviceRepo: repo,
	}
}

func (useCase *deviceUseCase) GetAll(c context.Context) ([]domains.Device, error) {
	res, err := useCase.deviceRepo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}
