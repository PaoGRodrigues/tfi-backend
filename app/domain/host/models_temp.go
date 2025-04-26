package host

// *********** Use Cases
// HostUseCase needs to be implemented in Host use cases
type HostUseCase interface {
	GetAllHosts() ([]Host, error)
	GetHosts() []Host
}

type HostBlocker interface {
	Block(string) (*string, error)
}

type HostsStorage interface {
	StoreHosts() error
	GetHostByIp(string) (Host, error)
}

// *********** Services
type HostService interface {
	GetAllHosts() ([]Host, error)
}

type HostBlockerService interface {
	BlockHost(string) error
}

// *********** Repositories
type HostsRepository interface {
	StoreHosts([]Host) error
	GetHostByIp(string) (Host, error)
}
