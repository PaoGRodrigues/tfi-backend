package traffic

import "net"

func FilterPublicServers(flows []Server) []Server {
	servers := []Server{}

	for _, srv := range flows {
		ip := net.ParseIP(srv.IP)
		if !ip.IsPrivate() {
			servers = append(servers, srv)
		}
	}
	return servers
}
