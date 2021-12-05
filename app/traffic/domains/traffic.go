package domains

import "time"

type Traffic struct {
	ID          int       `json:"ID"`
	Timestamp   time.Time `json:Timestamp`
	Source      string    `json:SourceIp`
	Destination string    `json:Destination`
	Port        string    `json:Port`
	Service     string    `json:Service`
	Bytes       int       `json:Bytes`
}

type TrafficUseCase interface {
	GetAllTraffic() ([]Traffic, error)
}

type TrafficRepository interface {
	GetAll() ([]Traffic, error)
}
