package usecase

import (
	"time"

	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	alertPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/alert"
)

type GetAlertsUseCase struct {
	repository alertPorts.AlertReader
}

func NewGetAlertsUseCase(repository alertPorts.AlertReader) *GetAlertsUseCase {
	return &GetAlertsUseCase{
		repository: repository,
	}
}

func (usecase *GetAlertsUseCase) GetAllAlerts() ([]alert.Alert, error) {
	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix()) //To get 7 days back

	alerts, err := usecase.repository.GetAllAlerts(epoch_begin, epoch_end)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}
