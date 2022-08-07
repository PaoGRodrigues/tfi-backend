package services

import (
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
)

type AlersFakeClient struct {
}

func NewAlertsFakeClient() *AlersFakeClient {

	return &AlersFakeClient{}
}

func (d *AlersFakeClient) GetAllAlerts() ([]domains.Alert, error) {

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
