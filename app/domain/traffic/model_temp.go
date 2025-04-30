package traffic

type BytesPerDestination struct {
	Bytes       int
	Destination string
}
type BytesPerCountry struct {
	Bytes   int    `json:"bytes"`
	Country string `json:"country"`
}

type TrafficBytesParser interface {
	GetBytesPerDestination() ([]BytesPerDestination, error)
	GetBytesPerCountry() ([]BytesPerCountry, error)
}

type TrafficStorage interface {
	StoreFlows() error
}
