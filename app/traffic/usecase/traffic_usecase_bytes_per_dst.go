package usecase

import (
	host_domains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type BytesAggregatorParser struct {
	trafficSearcher domains.TrafficUseCase
	hostsFilter     host_domains.HostsFilter
}

func NewBytesDestinationParser(trafSearcher domains.TrafficUseCase, hostsFilter host_domains.HostsFilter) *BytesAggregatorParser {
	return &BytesAggregatorParser{
		trafficSearcher: trafSearcher,
		hostsFilter:     hostsFilter,
	}
}

func (parser *BytesAggregatorParser) GetBytesPerDestination() ([]domains.BytesPerDestination, error) {
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

	parsedBytesDst := parsePerDest(activeFlows, remoteHosts)
	bytesDst := sumBytes(parsedBytesDst)
	return bytesDst, nil
}

func parsePerDest(actFlows []domains.ActiveFlow, remoteHosts []host_domains.Host) []domains.BytesPerDestination {
	bytesDst := []domains.BytesPerDestination{}

	for _, flow := range actFlows {
		var serverName string
		if flow.Server.Name != "" {
			serverName = flow.Server.Name
		} else {
			serverName = flow.Server.IP
		}
		for _, remh := range remoteHosts {
			if serverName == remh.Name || serverName == remh.IP {
				bpd := domains.BytesPerDestination{
					Bytes:       flow.Bytes,
					Destination: serverName,
					City:        remh.City,
					Country:     remh.Country,
				}
				bytesDst = append(bytesDst, bpd)
			}
		}
	}
	return bytesDst
}

func sumBytes(bpd []domains.BytesPerDestination) []domains.BytesPerDestination {
	m := map[string]int{}
	for _, v := range bpd {
		m[v.Destination] += v.Bytes
	}

	newBpd := []domains.BytesPerDestination{}
	for dest, bytes := range m {
		new := domains.BytesPerDestination{}
		new.Destination = dest
		new.Bytes = bytes
		for _, b := range bpd {
			if b.Destination == dest {
				new.City = b.City
				new.Country = b.Country
				break
			}
		}
		newBpd = append(newBpd, new)
	}

	return newBpd
}

func (parser *BytesAggregatorParser) GetBytesPerCountry() ([]domains.BytesPerCountry, error) {
	activeFlows := parser.trafficSearcher.GetActiveFlows()
	if len(activeFlows) == 0 {
		current, err := parser.trafficSearcher.GetAllActiveTraffic()
		if err != nil {
			return []domains.BytesPerCountry{}, err
		}
		activeFlows = current
	}

	remoteHosts, err := parser.hostsFilter.GetRemoteHosts()
	if err != nil {
		return []domains.BytesPerCountry{}, err
	}

	bytesDst := parsePerCountry(activeFlows, remoteHosts)
	return bytesDst, nil
}

type tempFlow struct {
	ip      string
	country string
	bytes   int
}

func parsePerCountry(actFlows []domains.ActiveFlow, remoteHosts []host_domains.Host) []domains.BytesPerCountry {

	det := []tempFlow{}
	for _, flow := range actFlows {
		for _, remh := range remoteHosts {
			if flow.Server.IP == remh.IP {
				temp := tempFlow{
					ip:      flow.Server.IP,
					country: remh.Country,
					bytes:   flow.Bytes,
				}
				det = append(det, temp)
			}
		}
	}

	bytesDst := sumCountries(det)
	return bytesDst
}

func sumCountries(temp []tempFlow) []domains.BytesPerCountry {
	m := map[string]int{}
	for _, v := range temp {
		m[v.country] += v.bytes
	}

	bsCountry := []domains.BytesPerCountry{}
	for k, v := range m {
		bpc := domains.BytesPerCountry{
			Country: k,
			Bytes:   v,
		}
		bsCountry = append(bsCountry, bpc)
	}
	return bsCountry
}
