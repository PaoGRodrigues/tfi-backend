package traffic

type BytesPerDestination struct {
	Bytes       int
	Destination string
}
type BytesPerCountry struct {
	Bytes   int    `json:"bytes"`
	Country string `json:"country"`
}

// *********** Use Cases
type TrafficUseCase interface {
	GetAllActiveTraffic() ([]TrafficFlow, error)
	GetActiveFlows() []TrafficFlow
}

type TrafficBytesParser interface {
	GetBytesPerDestination() ([]BytesPerDestination, error)
	GetBytesPerCountry() ([]BytesPerCountry, error)
}

type TrafficStorage interface {
	StoreFlows() error
}

// *********** Repositories
type TrafficRepository interface {
	StoreFlows([]TrafficFlow) error
	GetServerByAttr(string) (Server, error)
	GetClients() ([]Client, error)
	GetServers() ([]Server, error)
	GetFlowByKey(string) (TrafficFlow, error)
}

// *********** Services
type TrafficService interface {
	GetAllActiveTraffic() ([]TrafficFlow, error)
}
