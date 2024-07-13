package domains

// *********** Entities
// Host connected to the network
type Host struct {
	Name        string
	ASname      string `json:"ASname"`
	PrivateHost bool
	IP          string
	Mac         string
	City        string
	Country     string
}

// *********** Use Cases
// HostUseCase needs to be implemented in Host use cases
type HostUseCase interface {
	GetAllHosts() ([]Host, error)
	GetHosts() []Host
}

type HostsFilter interface {
	GetLocalHosts() ([]Host, error)
	GetRemoteHosts() ([]Host, error)
	GetHost(string) (Host, error)
}

type HostBlocker interface {
	Block(string) (Host, error)
}

// *********** Services
type HostService interface {
	GetAllHosts() ([]Host, error)
}

type HostBlockerService interface {
	BlockHost(Host) error
}
