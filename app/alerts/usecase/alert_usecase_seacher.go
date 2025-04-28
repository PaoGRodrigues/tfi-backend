package usecase

import (
	"time"

	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
)

type AlertSearcher struct {
	alertService alert.AlertService
	alerts       []alert.Alert
}

func NewAlertSearcher(service alert.AlertService) *AlertSearcher {
	return &AlertSearcher{
		alertService: service,
	}
}

func (searcher *AlertSearcher) GetAllAlerts() ([]alert.Alert, error) {
	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix()) //To get 7 days back

	alerts, err := searcher.GetAllAlertsByTime(epoch_begin, epoch_end)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}

func (searcher *AlertSearcher) GetAllAlertsByTime(epochBegin int, epochEnd int) ([]alert.Alert, error) {

	alerts := []alert.Alert{}
	res, err := searcher.alertService.GetAllAlerts(epochBegin, epochEnd)
	if err != nil {
		return nil, err
	}
	if res != nil {
		alerts = append(alerts, res...)
	}

	searcher.alerts = alerts
	return alerts, nil
}
