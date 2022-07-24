package usecase

import (
	host_domains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type BytesDestinationParser struct {
	trafficSearcher domains.TrafficUseCase
	hostsFilter     host_domains.HostsFilter
}

func NewBytesDestinationParser(trafSearcher domains.TrafficUseCase, hostsFilter host_domains.HostsFilter) *BytesDestinationParser {
	return &BytesDestinationParser{
		trafficSearcher: trafSearcher,
		hostsFilter:     hostsFilter,
	}
}

func (parser *BytesDestinationParser) GetBytesPerDestination() ([]domains.BytesPerDestination, error) {
	activeFlows := parser.trafficSearcher.GetActiveFlows()
	if len(activeFlows) == 0 {
		current, err := parser.trafficSearcher.GetAllActiveTraffic()
		if err != nil {
			return []domains.BytesPerDestination{}, err
		}
		activeFlows = current
	}

	remoteHosts, err := parser.hostsFilter.GetRemoteHosts()
	if err != nil {
		return []domains.BytesPerDestination{}, err
	}

	bytesDst := parse(activeFlows, remoteHosts)
	return bytesDst, nil
}

func parse(actFlows []domains.ActiveFlow, remoteHosts []host_domains.Host) []domains.BytesPerDestination {
	bytesDst := []domains.BytesPerDestination{}

	for _, flow := range actFlows {
		for _, remh := range remoteHosts {
			if flow.Server.IP == remh.IP {
				bpd := domains.BytesPerDestination{
					Bytes:       flow.Bytes,
					Destination: flow.Server.Name,
					City:        remh.City,
					Country:     remh.Country,
				}
				bytesDst = append(bytesDst, bpd)
			}
		}
	}
	return bytesDst
}
