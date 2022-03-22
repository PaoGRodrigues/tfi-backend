package domains

// Device connected to the network
type Device struct {
	PrivateHost bool
	Country     string
	Name        string
	IP          string
	OsDetail    string
	Mac         string
}

//DeviceUseCase needs to be implemented in Device use cases
type DeviceUseCase interface {
	GetAllDevices() ([]Device, error)
}

type DeviceRepository interface {
	GetAll() ([]Device, error)
}
