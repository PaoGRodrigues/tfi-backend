package domains

import "time"

type Traffic struct {
	ID          int       `json:"ID"`
	Datetime    time.Time `json:Timestamp`
	Source      string    `json:SourceIp`
	Destination string    `json:Destination`
	Port        int       `json:Port`
	Protocol    string    `json:Protocol`
	Service     string    `json:Service`
	Bytes       int       `json:Bytes`
}

type TrafficUseCase interface {
	GetAllTraffic() ([]Traffic, error)
}

type TrafficRepository interface {
	GetAll() ([]Traffic, error)
}
