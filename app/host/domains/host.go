package domains

// Host connected to the network
type Host struct {
	Name        string
	PrivateHost bool
	IP          string
	Mac         string
	City        string `json:"-"`
	Country     string `json:"-"`
}

//HostUseCase needs to be implemented in Host use cases
type HostUseCase interface {
	GetAllHosts() ([]Host, error)
	GetHosts() []Host
}

type HostsFilter interface {
	GetLocalHosts() ([]Host, error)
	GetRemoteHosts() ([]Host, error)
}

type HostRepository interface {
	GetAll() ([]Host, error)
}
