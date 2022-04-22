package usecase

import (
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type TrafficSearcher struct {
	trafficRepo domains.TrafficRepository
}

func NewTrafficSearcher(repo domains.TrafficRepository) *TrafficSearcher {

	return &TrafficSearcher{
		trafficRepo: repo,
	}
}

func (gw *TrafficSearcher) GetAllActiveTraffic() ([]domains.Traffic, error) {
	res, err := gw.trafficRepo.GetAllActiveTraffic()
	if err != nil {
		return nil, err
	}
	return res, nil
}
