package alert_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	usecase "github.com/PaoGRodrigues/tfi-backend/app/usecase/alert"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	alertPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/alert"

	"go.uber.org/mock/gomock"
)

func TestSendMessageSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := epoch_end - 300

	mockService := mocks.NewMockNotifier(ctrl)
	mockRepository := alertPortsMock.NewMockAlertReader(ctrl)
	mockRepository.EXPECT().GetAllAlerts(epoch_begin, epoch_end).Return([]alert.Alert{expected[0], expected[1]}, nil)
	alerts := alert.ParseAlerts(expected)
	mockService.EXPECT().SendMessage(alerts[0]).Return(nil)
	mockService.EXPECT().SendMessage(alerts[1]).Return(nil)

	alertNotif := usecase.NewNotifyAlertsUseCase(mockService, mockRepository)
	err := alertNotif.SendAlertMessages()
	if err != nil {
		t.Error("Testing error")
	}
}

func TestSendMessageReturnErrorWhenCallGetAllAlertsByTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := epoch_end - 300

	mockService := mocks.NewMockNotifier(ctrl)
	mockRepository := alertPortsMock.NewMockAlertReader(ctrl)
	mockRepository.EXPECT().GetAllAlerts(epoch_begin, epoch_end).Return(nil, errors.New("No alerts available"))

	alertNotif := usecase.NewNotifyAlertsUseCase(mockService, mockRepository)
	err := alertNotif.SendAlertMessages()
	if err == nil {
		t.Error("It's an error!")
	}
}

func TestSendMessageReturnErrorSendingAMessageButContinueAnyway(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := epoch_end - 300

	mockService := mocks.NewMockNotifier(ctrl)
	mockRepository := alertPortsMock.NewMockAlertReader(ctrl)
	mockRepository.EXPECT().GetAllAlerts(epoch_begin, epoch_end).Return([]alert.Alert{expected[0], expected[1]}, nil)
	alerts := alert.ParseAlerts(expected)
	mockService.EXPECT().SendMessage(alerts[0]).Return(fmt.Errorf("test error"))
	mockService.EXPECT().SendMessage(alerts[1]).Return(nil)

	alertNotif := usecase.NewNotifyAlertsUseCase(mockService, mockRepository)
	err := alertNotif.SendAlertMessages()
	if err != nil {
		t.Error("Testing error")
	}
}

func TestSendMessageReturnErrorWhenGetAllAlertsByTimeReturnZeroAlerts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockNotifier(ctrl)
	mockRepository := alertPortsMock.NewMockAlertReader(ctrl)
	mockRepository.EXPECT().GetAllAlerts(gomock.Any(), gomock.Any()).Return(nil, nil)

	alertNotif := usecase.NewNotifyAlertsUseCase(mockService, mockRepository)
	err := alertNotif.SendAlertMessages()
	if err == nil {
		t.Error("Testing error")
	}
}
