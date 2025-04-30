package traffic

import traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"

type TrafficDBRepository interface {
	StoreFlows([]traffic.TrafficFlow) error
	GetServerByAttr(string) (traffic.Server, error)
	GetClients() ([]traffic.Client, error)
	GetServers() ([]traffic.Server, error)
	GetFlowByKey(string) (traffic.TrafficFlow, error)
}
