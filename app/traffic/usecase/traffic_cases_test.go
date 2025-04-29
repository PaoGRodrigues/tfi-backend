package usecase_test

import (
	hosts "github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
)

var server1 = traffic.Server{
	IP:                "8.8.8.8",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "google.com.ar",
	Country:           "US",
	Key:               "12344567",
}

var server2 = traffic.Server{
	IP:                "8.8.8.8",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "google.com.ar",
	Country:           "US",
	Key:               "12344568",
}

var server3 = traffic.Server{
	IP:                "8.8.10.8",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "telegram.com",
	Country:           "RU",
	Key:               "12344569",
}

var noNameServer = traffic.Server{
	IP:                "8.8.10.10",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "",
	Country:           "US",
	Key:               "12344570",
}

var expectedFlowFromSearcher = []traffic.ActiveFlow{
	{
		Client: traffic.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: traffic.Server{
			IP:                "8.8.8.8",
			IsBroadcastDomain: false,
			IsDHCP:            false,
			Port:              443,
			Name:              "google.com.ar",
			Key:               "12344567",
			Country:           "US",
		},
		Protocol: traffic.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
}

var expectedFlowFromSearcherWithoutName = []traffic.ActiveFlow{
	{
		Client: traffic.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: noNameServer,
		Protocol: traffic.Protocol{
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
	{
		Name:        "",
		PrivateHost: false,
		IP:          "8.8.10.10",
		Country:     "US",
	},
}

var secondExpectedFlowFromSearcher = []traffic.ActiveFlow{
	{
		Client: traffic.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: server1,
		Protocol: traffic.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
	{
		Client: traffic.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: server2,
		Protocol: traffic.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 100000,
	},
}

var expectedPerCountrySearcher = []traffic.ActiveFlow{
	{
		Client: traffic.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: server1,
		Protocol: traffic.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
	{
		Client: traffic.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: server3,
		Protocol: traffic.Protocol{
			L4: "TCP",
			L7: "TLS.Telegram",
		},
		Bytes: 5566778,
	},
	{
		Client: traffic.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: traffic.Server{
			IP:                "8.8.10.82",
			IsBroadcastDomain: false,
			IsDHCP:            false,
			Port:              443,
			Name:              "telegram.com",
		},
		Protocol: traffic.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	},
}
