package services

import (
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
)

func (d *FakeTool) GetAllAlerts(epoch_begin, epoch_end int) ([]domains.Alert, error) {

	alerts := []domains.Alert{
		domains.Alert{
			Name:      "test",
			Subtype:   "network",
			Family:    "network",
			Timestamp: time.Time{},
			Score:     "1",
			Severity:  "2",
			Msg:       "testing Msg",
		},
		domains.Alert{
			Name:      "test1",
			Subtype:   "network",
			Family:    "network",
			Timestamp: time.Time{},
			Score:     "1",
			Severity:  "2",
			Msg:       "testing Msg",
		},
	}

	return alerts, nil
}
