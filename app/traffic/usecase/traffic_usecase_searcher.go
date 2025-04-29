package usecase

import (
	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
)

type TrafficSearcher struct {
	trafficService domains.TrafficService
	activeFlows    []domains.ActiveFlow
}

func NewTrafficSearcher(trafSrv domains.TrafficService) *TrafficSearcher {

	return &TrafficSearcher{
		trafficService: trafSrv,
	}
}

func (gw *TrafficSearcher) GetAllActiveTraffic() ([]domains.ActiveFlow, error) {
	res, err := gw.trafficService.GetAllActiveTraffic()
	if err != nil {
		return []domains.ActiveFlow{}, err
	}
	gw.activeFlows = res
	return res, nil
}

func (gw *TrafficSearcher) GetActiveFlows() []domains.ActiveFlow {
	return gw.activeFlows
}
