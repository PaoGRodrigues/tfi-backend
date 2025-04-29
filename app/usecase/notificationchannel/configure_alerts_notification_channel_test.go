package notificationchannel_test

import (
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
