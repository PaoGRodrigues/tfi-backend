package usecase_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/alerts/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	"github.com/golang/mock/gomock"
)

func TestGetAllAlertsReturnListOfAlerts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.Alert{
		domains.Alert{
			Name:      "test",
			Subtype:   "network",
			Family:    "network",
			Timestamp: time.Time{},
			Score:     "1",
			Severity:  "2",
			Msg:       "testing Msg",
		},
	}

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

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetAllAlertsReturnErrorWhenCallService(t *testing.T) {
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
