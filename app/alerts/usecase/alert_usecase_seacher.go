package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"

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
	res, err := searcher.alertService.GetAllHosts()
	if err != nil {
		return []domains.Alert{}, err
	}
	searcher.alerts = res
	return res, nil
}
