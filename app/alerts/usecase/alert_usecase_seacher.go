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

	res, err := searcher.alertService.GetAllAlerts(epoch_begin, epoch_end)
	if err != nil {
		return []domains.Alert{}, err
	}
	searcher.alerts = res
	return res, nil
}
