package usecase

import (
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	host_domains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
)

type AlertSearcher struct {
	alertService domains.AlertService
	alerts       []domains.Alert
	hostsFilter  host_domains.HostsFilter
}

func NewAlertSearcher(service domains.AlertService, hostsFilter host_domains.HostsFilter) *AlertSearcher {
	return &AlertSearcher{
		alertService: service,
		hostsFilter:  hostsFilter,
	}
}

func (searcher *AlertSearcher) GetAllAlerts() ([]domains.Alert, error) {
	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix()) //To get 7 days back

	hosts, err := searcher.hostsFilter.GetLocalHosts()
	if err != nil {
		return []domains.Alert{}, err
	}

	alerts := []domains.Alert{}
	for _, host := range hosts {
		res, err := searcher.alertService.GetAllAlerts(epoch_begin, epoch_end, host.IP)
		if err != nil {
			return []domains.Alert{}, err
		}
		alerts = append(alerts, res...)
	}
	searcher.alerts = alerts

	return alerts, nil
}
