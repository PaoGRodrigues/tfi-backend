package traffic

import (
	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
)

type TrafficSearcher struct {
	trafficService domains.TrafficService
	activeFlows    []domains.TrafficFlow
}

func NewTrafficSearcher(trafSrv domains.TrafficService) *TrafficSearcher {

	return &TrafficSearcher{
		trafficService: trafSrv,
	}
}

func (gw *TrafficSearcher) GetAllActiveTraffic() ([]domains.TrafficFlow, error) {
	res, err := gw.trafficService.GetAllActiveTraffic()
	if err != nil {
		return []domains.TrafficFlow{}, err
	}
	gw.activeFlows = res
	return res, nil
}

func (gw *TrafficSearcher) GetActiveFlows() []domains.TrafficFlow {
	return gw.activeFlows
}
