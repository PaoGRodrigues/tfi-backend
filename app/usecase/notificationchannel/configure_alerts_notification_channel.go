package notificationchannel

import (
	"errors"

	notificationChannelPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/notificationchannel"
)

type ConfigureChannelUseCase struct {
	notificationChannel notificationChannelPorts.NotificationChannel
}

func NewConfigureChannelUseCase(notificationChannel notificationChannelPorts.NotificationChannel) *ConfigureChannelUseCase {
	return &ConfigureChannelUseCase{
		notificationChannel: notificationChannel,
	}
}

func (usecase *ConfigureChannelUseCase) Configure(user string, token string) error {

	if user == "" {
		return errors.New("user is empty")
	}
	if token == "" {
		return errors.New("token is empty")
	}
	err := usecase.notificationChannel.Configure(user, token)
	if err != nil {
		return err
	}
	return nil
}
