package domains

// Device connected to the network
type Device struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	IP      string `json:"IP"`
	Details string `json:"Details"`
}

//DeviceUseCase needs to be implemented in Device use cases
type DeviceGateway interface {
	GetAll() ([]Device, error)
}

type DeviceRepository interface {
	GetAll() ([]Device, error)
}
