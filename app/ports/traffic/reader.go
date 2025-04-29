package traffic

import "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"

type TrafficReader interface {
	GetTrafficFlows() ([]traffic.TrafficFlow, error)
}
