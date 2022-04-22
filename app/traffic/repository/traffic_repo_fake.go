package repository

import (
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type TrafficFakeClient struct {
}

func NewTrafficFakeClient() *TrafficFakeClient {

	return &TrafficFakeClient{}
}

func (trff *TrafficFakeClient) GetAllActiveTraffic() ([]domains.Traffic, error) {
	trafficStruct := []domains.Traffic{
		domains.Traffic{
			ID:          1234,
			Datetime:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			Source:      "192.168.4.9",
			Destination: "lib.gen.rus",
			Port:        443,
			Protocol:    "tcp",
			Service:     "SSL",
			Bytes:       345,
		},
		domains.Traffic{
			ID:          1234,
			Datetime:    time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
			Source:      "192.168.4.9",
			Destination: "lib.gen.rus",
			Port:        443,
			Protocol:    "tcp",
			Service:     "SSL",
			Bytes:       10000,
		},
		domains.Traffic{
			ID:          1234,
			Datetime:    time.Date(2021, time.February, 20, 0, 0, 0, 0, time.UTC),
			Source:      "192.168.4.9",
			Destination: "lib.gen.rus",
			Port:        443,
			Protocol:    "tcp",
			Service:     "SSL",
			Bytes:       1000,
		},
	}

	return trafficStruct, nil
}
