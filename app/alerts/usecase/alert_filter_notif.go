package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
)

const minutes = 60

type AlertNotifier struct {
	searcher     domains.AlertUseCase
	notifService domains.Notifier
}

func NewAlertNotifier(service domains.Notifier, searcher domains.AlertUseCase) *AlertNotifier {
	return &AlertNotifier{
		searcher:     searcher,
		notifService: service,
	}
}

func (an *AlertNotifier) SendLastAlertMessages() error {
	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := epoch_end - minutes

	lastAlerts, err := an.searcher.GetAllAlertsByTime(epoch_begin, epoch_end)
	if err != nil {
		return err
	}
	if len(lastAlerts) == 0 {
		return errors.New("No alerts available")
	}
	parsedAlerts := ParseAlerts(lastAlerts)

	for _, alert := range parsedAlerts {
		err := an.notifService.SendMessage(alert)
		if err != nil {
			return errors.New("Cannot send message")
		}
	}
	return nil
}

func ParseAlerts(alerts []domains.Alert) []string {
	messages := []string{}

	for _, alert := range alerts {
		b, err := json.Marshal(alert)
		if err != nil {
			fmt.Println(err)
		}
		messages = append(messages, string(b))
	}
	return messages

}
