package domains

// Device connected to the network
type Device struct {
	IsLocalhost bool   `json:is_localhost"`
	Country     string `json:country"`
	Name        string `json:name"`
	IP          string `json:ip"`
	OsDetail    string `json:"os_detail"`
	Mac         string `json:mac"`
}

//DeviceUseCase needs to be implemented in Device use cases
type DeviceUseCase interface {
	GetAllDevices() ([]Device, error)
}

type DeviceRepository interface {
	GetAll() ([]Device, error)
}
