package alert

// *********** Use Cases
// AlertUseCase needs to be implemented in Alert use cases
type AlertUseCase interface {
	GetAllAlerts() ([]Alert, error)
	GetAllAlertsByTime(int, int) ([]Alert, error)
}

type AlertsSender interface {
	SendLastAlertMessages() error
}

type Notifier interface {
	SendMessage(string) error
}
