package domains

import "time"

//*********** Entities
// Alerts
type Alert struct {
	Name      string
	Subtype   string
	Family    string
	Timestamp time.Time
	Score     string
	Severity  string
	Msg       string
}

//*********** Use Cases
//AlertUseCase needs to be implemented in Alert use cases
type AlertUseCase interface {
	GetAllAlerts() ([]Alert, error)
}

//*********** Services
type AlertService interface {
	GetAllAlerts(int, int) ([]Alert, error)
}
