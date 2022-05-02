package domains

type ActiveFlow struct {
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

type TrafficUseCase interface {
	GetAllActiveTraffic() ([]ActiveFlow, error)
	GetActiveFlows() []ActiveFlow
}

type TrafficRepository interface {
	GetAllActiveTraffic() ([]ActiveFlow, error)
}
