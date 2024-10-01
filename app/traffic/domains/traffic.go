package domains

// *********** Entities
type ActiveFlow struct {
	Key       string `json:"key"`
	FirstSeen uint64 `json:"first_seen"` //Unix timestamp
	LastSeen  uint64 `json:"last_seen"`  //Unix timestamp
	Client    Client
	Server    Server
	Bytes     int
	Protocol  Protocol `json:",omitempty"`
}

type Client struct {
	Key  string `json:",omitempty"`
	Name string `json:",omitempty"`
	Port int
	IP   string
}

type Server struct {
	Key               string `json:",omitempty"`
	IP                string
	IsBroadcastDomain bool `json:",omitempty"`
	IsDHCP            bool `json:",omitempty"`
	Port              int
	Name              string `json:",omitempty"`
	Country           string `json:",omitempty"`
}

type Protocol struct {
	Key   string `json:",omitempty"`
	L4    string `json:",omitempty"`
	L7    string `json:",omitempty"`
	Label string `json:",omitempty"`
}

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
	GetAllActiveTraffic() ([]ActiveFlow, error)
	GetActiveFlows() []ActiveFlow
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
	StoreFlows([]ActiveFlow) error
	GetServerByAttr(string) (Server, error)
	GetClients() ([]Client, error)
	GetServers() ([]Server, error)
	GetFlowByKey(string) (ActiveFlow, error)
}

// *********** Services
type TrafficService interface {
	GetAllActiveTraffic() ([]ActiveFlow, error)
}
