package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	alertPort "github.com/PaoGRodrigues/tfi-backend/app/ports/alert"
)

const seconds = 300
const cybersecurity = "Cybersecurity"

type AlertNotifier struct {
	repository   alertPort.AlertReader
	notifService alert.Notifier
}

func NewAlertNotifier(service alert.Notifier, repository alertPort.AlertReader) *AlertNotifier {
	return &AlertNotifier{
		repository:   repository,
		notifService: service,
	}
}

func (an *AlertNotifier) SendLastAlertMessages() error {
	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := epoch_end - seconds

	lastAlerts, err := an.repository.GetAllAlerts(epoch_begin, epoch_end)
	if err != nil {
		return err
	}
	if lastAlerts == nil {
		return errors.New("No alerts available")
	}
	parsedAlerts := ParseAlerts(lastAlerts)

	for _, alert := range parsedAlerts {
		err := an.notifService.SendMessage(alert)
		if err != nil {
			//It won't stop if a message can't be sent
			fmt.Printf("Cannot send message")
		}
	}
	return nil
}

func ParseAlerts(alerts []alert.Alert) []string {
	messages := []string{}

	for _, alert := range alerts {
		if alert.Category == cybersecurity {
			b, err := json.Marshal(alert)
			if err != nil {
				fmt.Println(err)
			}
			messages = append(messages, string(b))
		}
	}
	return messages

}
