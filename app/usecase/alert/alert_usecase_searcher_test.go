package usecase_test

import (
	"fmt"
	"testing"
	"time"

	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	usecase "github.com/PaoGRodrigues/tfi-backend/app/usecase/alert"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func TestGetAllAlertsReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix())

	mockService := mocks.NewMockAlertService(ctrl)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end).Return([]alert.Alert{}, fmt.Errorf("test error"))

	alertSearcher := usecase.NewAlertSearcher(mockService)
	_, err := alertSearcher.GetAllAlerts()

	if err == nil {
		t.Fail()
	}
}

func TestGetAllAlertsReturnAlerts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix())

	mockService := mocks.NewMockAlertService(ctrl)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end).Return(expected, nil)

	alertSearcher := usecase.NewAlertSearcher(mockService)
	got, err := alertSearcher.GetAllAlerts()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected, got)
}
