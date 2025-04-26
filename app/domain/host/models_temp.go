package host

// *********** Repositories
type HostsRepository interface {
	StoreHosts([]Host) error
	GetHostByIp(string) (Host, error)
}
