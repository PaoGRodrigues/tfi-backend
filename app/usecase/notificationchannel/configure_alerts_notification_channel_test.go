package notificationchannel_test

import (
	"fmt"
	"testing"

	usecase "github.com/PaoGRodrigues/tfi-backend/app/usecase/notificationchannel"
	notificationChannelPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/notificationchannel"

	"go.uber.org/mock/gomock"
)

func TestConfigureChannelSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChannel := notificationChannelPortsMock.NewMockNotificationChannel(ctrl)
	mockChannel.EXPECT().Configure("user", "token").Return(nil)

	configureNotificationChannelUseCase := usecase.NewConfigureChannelUseCase(mockChannel)
	err := configureNotificationChannelUseCase.Configure("user", "token")

	if err != nil {
		t.Error("Testing error")
	}
}

func TestConfigureChannelReturnErrorBecauseUserIsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChannel := notificationChannelPortsMock.NewMockNotificationChannel(ctrl)

	configureNotificationChannelUseCase := usecase.NewConfigureChannelUseCase(mockChannel)
	err := configureNotificationChannelUseCase.Configure("", "token")

	if err == nil {
		t.Error("Testing error")
	}
}

func TestConfigureChannelReturnErrorBecauseTokenIsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChannel := notificationChannelPortsMock.NewMockNotificationChannel(ctrl)

	configureNotificationChannelUseCase := usecase.NewConfigureChannelUseCase(mockChannel)
	err := configureNotificationChannelUseCase.Configure("user", "")

	if err == nil {
		t.Error("Testing error")
	}
}

func TestConfigureChannelReturnErrorBecauseChannelReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChannel := notificationChannelPortsMock.NewMockNotificationChannel(ctrl)
	mockChannel.EXPECT().Configure("user", "token").Return(fmt.Errorf("Error test"))

	configureNotificationChannelUseCase := usecase.NewConfigureChannelUseCase(mockChannel)
	err := configureNotificationChannelUseCase.Configure("user", "token")

	if err == nil {
		t.Error("Testing error")
	}
}
