package host

// *********** Use Cases

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
