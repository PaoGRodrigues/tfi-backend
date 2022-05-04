package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"

type BytesDestinationParser struct {
	searcher domains.TrafficUseCase
}

func NewBytesDestinationParser(trafSearcher domains.TrafficUseCase) *BytesDestinationParser {
	return &BytesDestinationParser{
		searcher: trafSearcher,
	}
}

func (parser *BytesDestinationParser) GetBytesPerDestination() ([]domains.BytesPerDestination, error) {
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
