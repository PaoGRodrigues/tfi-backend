package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	"github.com/golang/mock/gomock"
)

func TestSendMessageSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := epoch_end - 60

	mockService := mocks.NewMockNotifier(ctrl)
	mockSearcher := mocks.NewMockAlertUseCase(ctrl)
	mockSearcher.EXPECT().GetAllAlertsByTime(epoch_begin, epoch_end).Return(expected, nil)
	alerts := usecase.ParseAlerts(expected)
	mockService.EXPECT().SendMessage(alerts[0]).Return(nil)
	mockService.EXPECT().SendMessage(alerts[1]).Return(nil)

	alertNotif := usecase.NewAlertNotifier(mockService, mockSearcher)
	err := alertNotif.SendLastAlertMessages()
	if err != nil {
		t.Error("Testing error")
	}
}

func TestSendMessageReturnErrorWhenCallGetAllAlertsByTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := epoch_end - 60

	mockService := mocks.NewMockNotifier(ctrl)
	mockSearcher := mocks.NewMockAlertUseCase(ctrl)
	mockSearcher.EXPECT().GetAllAlertsByTime(epoch_begin, epoch_end).Return(nil, errors.New("No alerts available"))

	alertNotif := usecase.NewAlertNotifier(mockService, mockSearcher)
	err := alertNotif.SendLastAlertMessages()
	if err == nil {
		t.Error("Testing error")
	}
}
