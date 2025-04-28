package usecase

import (
	"time"

	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	alertPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/alert"
)

type GetAlertsUseCase struct {
	repository alertPorts.AlertReader
	alerts     []alert.Alert
}

func NewGetAlertsUseCase(service alertPorts.AlertReader) *GetAlertsUseCase {
	return &GetAlertsUseCase{
		repository: service,
	}
}

func (searcher *GetAlertsUseCase) GetAllAlerts() ([]alert.Alert, error) {
	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix()) //To get 7 days back

	alerts, err := searcher.GetAllAlertsByTime(epoch_begin, epoch_end)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}

func (searcher *GetAlertsUseCase) GetAllAlertsByTime(epochBegin int, epochEnd int) ([]alert.Alert, error) {

	alerts := []alert.Alert{}
	res, err := searcher.repository.GetAllAlerts(epochBegin, epochEnd)
	if err != nil {
		return nil, err
	}
	if res != nil {
		alerts = append(alerts, res...)
	}

	searcher.alerts = alerts
	return alerts, nil
}
