package alert

// *********** Use Cases
// AlertReaderTemp needs to be implemented in Alert use cases
type AlertReaderTemp interface {
	GetAllAlerts() ([]Alert, error)
	GetAllAlertsByTime(int, int) ([]Alert, error)
}

type AlertsSender interface {
	SendLastAlertMessages() error
}

type Notifier interface {
	SendMessage(string) error
}
