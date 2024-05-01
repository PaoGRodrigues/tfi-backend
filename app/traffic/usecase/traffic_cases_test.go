package usecase_test

import (
	hosts "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

var server1 = domains.Server{
	IP:                "8.8.8.8",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "google.com.ar",
	Country:           "US",
	Key:               "12344567",
}

var server2 = domains.Server{
	IP:                "8.8.8.8",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "google.com.ar",
	Country:           "US",
	Key:               "12344568",
}

var server3 = domains.Server{
	IP:                "8.8.10.8",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "telegram.com",
	Country:           "RU",
	Key:               "12344569",
}

var expectedFlowFromSearcher = []domains.ActiveFlow{
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
			Key:               "12344567",
			Country:           "US",
		},
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
}

var expectedFlowFromSearcherWithoutName = []domains.ActiveFlow{
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
			Name:              "",
			Country:           "US",
		},
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
}

var expectedHosts = []hosts.Host{
	{
		Name:        "google.com.ar",
		PrivateHost: false,
		IP:          "8.8.8.8",
		Country:     "US",
		City:        "California",
	},
	{
		Name:        "sarasa2",
		PrivateHost: false,
		IP:          "198.8.8.8",
		Country:     "AR",
	},
	{
		Name:        "telegram.com",
		PrivateHost: false,
		IP:          "8.8.10.8",
		Country:     "RU",
	},
}

var secondExpectedFlowFromSearcher = []domains.ActiveFlow{
	{
		Client: domains.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: server1,
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
		Server: server2,
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 100000,
	},
}

var expectedPerCountrySearcher = []domains.ActiveFlow{
	{
		Client: domains.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: server1,
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
		Server: server3,
		Protocol: domains.Protocol{
			L4: "TCP",
			L7: "TLS.Telegram",
		},
		Bytes: 5566778,
	},
}
