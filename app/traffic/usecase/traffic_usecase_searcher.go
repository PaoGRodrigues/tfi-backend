package usecase

import (
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type TrafficSearcher struct {
	trafficRepo domains.TrafficRepository
	activeFlows []domains.ActiveFlow
}

func NewTrafficSearcher(repo domains.TrafficRepository) *TrafficSearcher {

	return &TrafficSearcher{
		trafficRepo: repo,
	}
}

func (gw *TrafficSearcher) GetAllActiveTraffic() ([]domains.ActiveFlow, error) {
	res, err := gw.trafficRepo.GetAllActiveTraffic()
	if err != nil {
		return []domains.ActiveFlow{}, err
	}
	gw.activeFlows = res
	return res, nil
}

func (gw *TrafficSearcher) GetActiveFlows() []domains.ActiveFlow {
	return gw.activeFlows
}
