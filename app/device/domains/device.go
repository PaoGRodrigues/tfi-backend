package domains

// Device connected to the network
type Device struct {
	ID      int    `json:"-"`
	Name    string `json:"name"`
	IP      string `json:"ip"`
	Details string `json:"details"`
}

//DeviceUseCase needs to be implemented in Device use cases
type DeviceGateway interface {
	GetAll() ([]Device, error)
}

type DeviceRepository interface {
	GetAll() ([]Device, error)
}
