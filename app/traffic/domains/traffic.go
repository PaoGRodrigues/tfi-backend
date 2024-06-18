package domains

// *********** Entities
type ActiveFlow struct {
	Key       string `json:"key"`
	FirstSeen uint64 `json:"first_seen"` //Unix timestamp
	LastSeen  uint64 `json:"last_seen"`  //Unix timestamp
	Client    Client
	Server    Server
	Bytes     int
	Protocol  Protocol
}

type Client struct {
	Key  string
	Name string
	Port int
	IP   string
}

type Server struct {
	Key               string
	IP                string
	IsBroadcastDomain bool
	IsDHCP            bool
	Port              int
	Name              string
	Country           string
}

type Protocol struct {
	Key   string
	L4    string
	L7    string
	Label string
}

type BytesPerDestination struct {
	Bytes       int
	Destination string
	Country     string
}
type BytesPerCountry struct {
	Bytes   int    `json:"bytes"`
	Country string `json:"country"`
}

// *********** Use Cases
type TrafficUseCase interface {
	GetAllActiveTraffic() ([]ActiveFlow, error)
	GetActiveFlows() []ActiveFlow
}

type TrafficActiveFlowsSearcher interface {
	GetBytesPerDestination() ([]BytesPerDestination, error)
	GetBytesPerCountry() ([]BytesPerCountry, error)
}

type TrafficStorage interface {
	StoreFlows() error
}

// *********** Services
type TrafficService interface {
	GetAllActiveTraffic() ([]ActiveFlow, error)
}

// *********** Repositories
type TrafficRepository interface {
	StoreFlows([]ActiveFlow) error
	GetServerByAttr(string) (Server, error)
	GetClients() ([]Client, error)
	GetServers() ([]Server, error)
	GetFlowByKey(string) (ActiveFlow, error)
}
