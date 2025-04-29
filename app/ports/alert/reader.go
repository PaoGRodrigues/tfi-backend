package alert

import "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"

type AlertReader interface {
	GetAllAlerts(int, int) ([]alert.Alert, error)
}
