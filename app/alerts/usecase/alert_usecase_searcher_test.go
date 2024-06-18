package usecase_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/alerts/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestGetAllAlertsReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix())

	mockService := mocks.NewMockAlertService(ctrl)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end).Return([]domains.Alert{}, fmt.Errorf("test error"))

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
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end).Return([]domains.Alert{expected[0], expected[1]}, nil)

	alertSearcher := usecase.NewAlertSearcher(mockService)
	got, err := alertSearcher.GetAllAlerts()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected, got)
}
