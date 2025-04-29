package usecase

import (
	"errors"
	"fmt"
	"time"

	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	alertPort "github.com/PaoGRodrigues/tfi-backend/app/ports/alert"
)

const seconds = 300
const cybersecurity = "Cybersecurity"

type NotifyAlertsUseCase struct {
	repository   alertPort.AlertReader
	notifService alertPort.Notifier
}

func NewNotifyAlertsUseCase(service alertPort.Notifier, repository alertPort.AlertReader) *NotifyAlertsUseCase {
	return &NotifyAlertsUseCase{
		repository:   repository,
		notifService: service,
	}
}

func (an *NotifyAlertsUseCase) SendLastAlertMessages() error {
	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := epoch_end - seconds

	lastAlerts, err := an.repository.GetAllAlerts(epoch_begin, epoch_end)
	if err != nil {
		return err
	}
	if len(lastAlerts) == 0 {
		return errors.New("No alerts available")
	}
	parsedAlerts := alert.ParseAlerts(lastAlerts)

	for _, alert := range parsedAlerts {
		err := an.notifService.SendMessage(alert)
		if err != nil {
			//It won't stop if a message can't be sent
			fmt.Printf("Cannot send message")
		}
	}
	return nil
}
