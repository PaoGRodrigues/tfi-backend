package usecase_test

import (
	hosts "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

var expectedFlowFromSearcher = []domains.ActiveFlow{
	domains.ActiveFlow{
		Client: domains.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: domains.Server{
			IP:                "8.8.8.8",
			IsBroadcastDomain: false,
			IsDHCP:            false,
			Port:              443,
			Name:              "google.com.ar",
		},
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
}

var expectedFlowFromSearcherWithoutName = []domains.ActiveFlow{
	domains.ActiveFlow{
		Client: domains.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: domains.Server{
			IP:                "8.8.8.8",
			IsBroadcastDomain: false,
			IsDHCP:            false,
			Port:              443,
			Name:              "",
		},
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
}

var expectedHosts = []hosts.Host{
	hosts.Host{
		Name:        "google.com.ar",
		PrivateHost: false,
		IP:          "8.8.8.8",
		Country:     "US",
		City:        "California",
	},
	hosts.Host{
		Name:        "sarasa2",
		PrivateHost: false,
		IP:          "198.8.8.8",
		City:        "AR",
	},
}

var secondExpectedFlowFromSearcher = []domains.ActiveFlow{
	{
		Client: domains.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: domains.Server{
			IP:                "8.8.8.8",
			IsBroadcastDomain: false,
			IsDHCP:            false,
			Port:              443,
			Name:              "google.com.ar",
		},
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
	{
		Client: domains.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: domains.Server{
			IP:                "8.8.8.8",
			IsBroadcastDomain: false,
			IsDHCP:            false,
			Port:              443,
			Name:              "google.com.ar",
		},
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 100000,
	},
}

var expectedPerCountrySearcher = []domains.ActiveFlow{
	domains.ActiveFlow{
		Client: domains.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: domains.Server{
			IP:                "8.8.8.8",
			IsBroadcastDomain: false,
			IsDHCP:            false,
			Port:              443,
			Name:              "google.com.ar",
		},
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
	domains.ActiveFlow{
		Client: domains.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: domains.Server{
			IP:                "8.8.10.8",
			IsBroadcastDomain: false,
			IsDHCP:            false,
			Port:              443,
			Name:              "telegram.com",
		},
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
}

var expectedHostsPerCountry = []hosts.Host{
	hosts.Host{
		Name:        "google.com.ar",
		PrivateHost: false,
		IP:          "8.8.8.8",
		Country:     "US",
		City:        "California",
	},
	hosts.Host{
		Name:        "sarasa2",
		PrivateHost: false,
		IP:          "198.8.8.8",
		City:        "AR",
	},
	hosts.Host{
		Name:        "telegram.com",
		PrivateHost: false,
		IP:          "8.8.10.8",
		Country:     "US",
	},
}
