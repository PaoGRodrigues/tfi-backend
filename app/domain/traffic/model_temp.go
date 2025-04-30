package traffic

type BytesPerCountry struct {
	Bytes   int    `json:"bytes"`
	Country string `json:"country"`
}

type TrafficBytesParser interface {
	GetBytesPerCountry() ([]BytesPerCountry, error)
}
