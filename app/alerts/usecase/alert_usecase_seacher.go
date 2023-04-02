package usecase

import (
	"fmt"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	traffic_domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type AlertSearcher struct {
	alertService  domains.AlertService
	alerts        []domains.Alert
	trafficFilter traffic_domains.ActiveFlowsStorage
}

func NewAlertSearcher(service domains.AlertService, filter traffic_domains.ActiveFlowsStorage) *AlertSearcher {
	return &AlertSearcher{
		alertService:  service,
		trafficFilter: filter,
	}
}

func (searcher *AlertSearcher) GetAllAlerts() ([]domains.Alert, error) {
	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix()) //To get 7 days back

	alerts := []domains.Alert{}

	host := traffic_domains.Client{}

	res, err := searcher.alertService.GetAllAlerts(epoch_begin, epoch_end, host.IP)
	fmt.Print(res)
	if err != nil {
		return []domains.Alert{}, err
	}
	alerts = append(alerts, res...)

	/*
		clients, err := searcher.trafficFilter.GetClientsList()
		if err != nil {
			return []domains.Alert{}, err
		}

		alerts := []domains.Alert{}
		for _, host := range clients {
			res, err := searcher.alertService.GetAllAlerts(epoch_begin, epoch_end, host.IP)
			fmt.Print(res)
			if err != nil {
				return []domains.Alert{}, err
			}
			alerts = append(alerts, res...)
		}

		servers, err := searcher.trafficFilter.GetServersList()
		if err != nil {
			return []domains.Alert{}, err
		}
		for _, host := range servers {
			res, err := searcher.alertService.GetAllAlerts(epoch_begin, epoch_end, host.IP)
			fmt.Print(res)
			if err != nil {
				return []domains.Alert{}, err
			}
			alerts = append(alerts, res...)
		}*/
	searcher.alerts = alerts

	return alerts, nil
}
