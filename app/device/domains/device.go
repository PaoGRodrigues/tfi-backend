package domains

// Device connected to the network
type Device struct {
	IsLocalhost bool `json:is_localhost"`
	Country     string
	Name        string
	IP          string
	OsDetail    string `json:"os_detail"`
	Mac         string
}

//DeviceUseCase needs to be implemented in Device use cases
type DeviceUseCase interface {
	GetAllDevices() ([]Device, error)
}

type DeviceRepository interface {
	GetAll() ([]Device, error)
}
