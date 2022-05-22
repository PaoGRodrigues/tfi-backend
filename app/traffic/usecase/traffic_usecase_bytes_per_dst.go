package usecase

import (
	host_domains "github.com/PaoGRodrigues/tfi-backend/app/host/domains"
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
	actF, err := parser.getRemoteServerActiveFlows(activeFlows)
	if err != nil {
		return []domains.BytesPerDestination{}, err
	}
	bytesDst := parse(actF)
	return bytesDst, nil
}

func parse(actFlows []domains.ActiveFlow) []domains.BytesPerDestination {
	bytesDst := []domains.BytesPerDestination{}

	for _, flow := range actFlows {
		bpd := domains.BytesPerDestination{
			Bytes:       flow.Bytes,
			Destination: flow.Server.Name,
		}
		bytesDst = append(bytesDst, bpd)
	}
	return bytesDst
}

func (parser *BytesDestinationParser) getRemoteServerActiveFlows(activeFlows []domains.ActiveFlow) ([]domains.ActiveFlow, error) {
	remoteServerActiveFlows := []domains.ActiveFlow{}
	remoteHosts, err := parser.hostsFilter.GetRemoteHosts()

	if err != nil {
		return []domains.ActiveFlow{}, err
	}
	for _, ac := range activeFlows {
		for _, remh := range remoteHosts {
			if ac.Server.IP == remh.IP {
				remoteServerActiveFlows = append(remoteServerActiveFlows, ac)
			}
		}
	}
	return remoteServerActiveFlows, nil
}
