package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"

type BytesDestinationParser struct {
	searcher domains.TrafficUseCase
}

type BytesPerDestination struct {
	Bytes       int
	Destination string
}

func NewBytesDestinationParser(trafSearcher domains.TrafficUseCase) *BytesDestinationParser {
	return &BytesDestinationParser{
		searcher: trafSearcher,
	}
}

func (parser *BytesDestinationParser) GetBytesPerDestination() ([]BytesPerDestination, error) {
	activeFlows := parser.searcher.GetActiveFlows()
	if len(activeFlows) == 0 {
		current, err := parser.searcher.GetAllActiveTraffic()
		if err != nil {
			return nil, err
		}
		activeFlows = current
	}
	bytesDst := parse(activeFlows)
	return bytesDst, nil
}

func parse(actFlows []domains.ActiveFlow) []BytesPerDestination {
	bytesDst := []BytesPerDestination{}

	for _, flow := range actFlows {
		bpd := BytesPerDestination{
			Bytes:       flow.Bytes,
			Destination: flow.Server.Name,
		}
		bytesDst = append(bytesDst, bpd)
	}
	return bytesDst
}
