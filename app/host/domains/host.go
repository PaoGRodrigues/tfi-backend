package domains

// Host connected to the network
type Host struct {
	Name        string
	PrivateHost bool
	IP          string
	Mac         string
	City        string
	Country     string
}

//HostUseCase needs to be implemented in Host use cases
type HostUseCase interface {
	GetAllHosts() ([]Host, error)
	GetHosts() []Host
}

type LocalHostFilter interface {
	GetLocalHosts() ([]Host, error)
}

type HostRepository interface {
	GetAll() ([]Host, error)
}
