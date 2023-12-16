package usecase

import (
	"net"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type BytesAggregatorParser struct {
	flowsStorage domains.ActiveFlowsStorage
}

func NewBytesParser(flowsStorage domains.ActiveFlowsStorage) *BytesAggregatorParser {
	return &BytesAggregatorParser{
		flowsStorage: flowsStorage,
	}
}

func (parser *BytesAggregatorParser) GetBytesPerDestination() ([]domains.BytesPerDestination, error) {
	serversList, err := parser.flowsStorage.GetServersList()

	if err != nil {
		return []domains.BytesPerDestination{}, err
	}

	servers := filterPublicServers(serversList)

	flows := []domains.ActiveFlow{}
	for _, server := range servers {
		flow, err := parser.flowsStorage.GetFlowByKey(server.Key)
		if err != nil {
			return []domains.BytesPerDestination{}, err
		}
		flows = append(flows, flow)
	}

	parsedBytesDst := parsePerDest(flows)
	bytesDst := sumBytes(parsedBytesDst)
	return bytesDst, nil
}

func filterPublicServers(flows []domains.Server) []domains.Server {
	servers := []domains.Server{}

	for _, srv := range flows {
		ip := net.ParseIP(srv.IP)
		if !ip.IsPrivate() {
			servers = append(servers, srv)
		}
	}
	return servers
}

func parsePerDest(flows []domains.ActiveFlow) []domains.BytesPerDestination {

	bytesDst := []domains.BytesPerDestination{}

	for _, flow := range flows {
		var serverName string
		if flow.Server.Name != "" {
			serverName = flow.Server.Name
		} else {
			serverName = flow.Server.IP
		}
		bpd := domains.BytesPerDestination{
			Bytes:       flow.Bytes,
			Destination: serverName,
			Country:     flow.Server.Country,
		}
		bytesDst = append(bytesDst, bpd)
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
				new.Country = b.Country
				break
			}
		}
		newBpd = append(newBpd, new)
	}

	return newBpd
}

func (parser *BytesAggregatorParser) GetBytesPerCountry() ([]domains.BytesPerCountry, error) {
	serversList, err := parser.flowsStorage.GetServersList()

	if err != nil {
		return []domains.BytesPerCountry{}, err
	}

	servers := filterPublicServers(serversList)

	flows := []domains.ActiveFlow{}
	for _, server := range servers {
		flow, err := parser.flowsStorage.GetFlowByKey(server.Key)
		if err != nil {
			return []domains.BytesPerCountry{}, err
		}
		flows = append(flows, flow)
	}

	bytesCountry := sumCountries(flows)
	return bytesCountry, nil
}

func sumCountries(flow []domains.ActiveFlow) []domains.BytesPerCountry {
	m := map[string]int{}
	for _, v := range flow {
		m[v.Server.Country] += v.Bytes
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
