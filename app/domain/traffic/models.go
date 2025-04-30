package traffic

// *********** Entities
type TrafficFlow struct {
	Key       string `json:"key"`
	FirstSeen uint64 `json:"first_seen"` //Unix timestamp
	LastSeen  uint64 `json:"last_seen"`  //Unix timestamp
	Client    Client
	Server    Server
	Bytes     int
	Protocol  Protocol `json:",omitempty"`
}

type Client struct {
	Key  string `json:",omitempty"`
	Name string `json:",omitempty"`
	Port int
	IP   string
}

type Server struct {
	Key               string `json:",omitempty"`
	IP                string
	IsBroadcastDomain bool `json:",omitempty"`
	IsDHCP            bool `json:",omitempty"`
	Port              int
	Name              string `json:",omitempty"`
	Country           string `json:",omitempty"`
}

type Protocol struct {
	Key   string `json:",omitempty"`
	L4    string `json:",omitempty"`
	L7    string `json:",omitempty"`
	Label string `json:",omitempty"`
}
