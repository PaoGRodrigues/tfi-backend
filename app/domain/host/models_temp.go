package host

// *********** Use Cases

type HostWriter interface {
	StoreHosts() error
	GetHostByIp(string) (Host, error)
}

// *********** Repositories
type HostsRepository interface {
	StoreHosts([]Host) error
	GetHostByIp(string) (Host, error)
}
