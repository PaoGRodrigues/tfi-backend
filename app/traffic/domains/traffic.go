package domains

//*********** Entities
type ActiveFlow struct {
	Key      string
	FistSeen uint64 //Unix timestamp
	LastSeen uint64 //Unix timestamp
	Client   Client
	Server   Server
	Bytes    int
	Protocol Protocol
}

type Client struct {
	Name string
	Port int
	IP   string
}

type Server struct {
	IP                string
	IsBroadcastDomain bool
	IsDHCP            bool
	Port              int
	Name              string
}

type Protocol struct {
	L4 string
	L7 string
}

type BytesPerDestination struct {
	Bytes       int
	Destination string
	City        string
	Country     string
}

//*********** Use Cases
type TrafficUseCase interface {
	GetAllActiveTraffic() ([]ActiveFlow, error)
	GetActiveFlows() []ActiveFlow
}

type TrafficActiveFlowsSearcher interface {
	GetBytesPerDestination() ([]BytesPerDestination, error)
}

type ActiveFlowsStorage interface {
	StoreFlows() error
}

//*********** Repositories
type TrafficRepoClient interface {
	GetAllActiveTraffic() ([]ActiveFlow, error)
}

type TrafficRepoStore interface {
	StoreActiveFlows([]ActiveFlow) error
}
