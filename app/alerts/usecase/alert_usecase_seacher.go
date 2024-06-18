package usecase

import (
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
)

type AlertSearcher struct {
	alertService domains.AlertService
	alerts       []domains.Alert
}

func NewAlertSearcher(service domains.AlertService) *AlertSearcher {
	return &AlertSearcher{
		alertService: service,
	}
}

func (searcher *AlertSearcher) GetAllAlerts() ([]domains.Alert, error) {
	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix()) //To get 7 days back

	alerts, err := searcher.GetAllAlertsByTime(epoch_begin, epoch_end)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}

func (searcher *AlertSearcher) GetAllAlertsByTime(epochBegin int, epochEnd int) ([]domains.Alert, error) {

	alerts := []domains.Alert{}
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
