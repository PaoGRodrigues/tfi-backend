package domains

// Device connected to the network
type Device struct {
	Name        string
	PrivateHost bool
	IP          string
	Mac         string
	City        string
	Country     string
}

//DeviceUseCase needs to be implemented in Device use cases
type DeviceUseCase interface {
	GetAllDevices() ([]Device, error)
}

type DeviceRepository interface {
	GetAll() ([]Device, error)
}
