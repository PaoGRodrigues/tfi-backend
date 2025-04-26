package host

func GetLocalHosts(current []Host) []Host {
	localHosts := []Host{}
	if len(current) != 0 {
		for _, host := range current {
			if host.PrivateHost {
				localHosts = append(localHosts, host)
			}
		}
	}
	return localHosts
}
