package usecase_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/alerts/usecase"
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	mocks_host "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	"github.com/golang/mock/gomock"
)

func TestGetAllAlertsReturnListOfAlerts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.Alert{
		domains.Alert{

			Name:      "test",
			Family:    "flow",
			Timestamp: time.Time{},
			Score:     "10",
			Severity:  domains.Severity{Label: "2"},
			AlertFlow: domains.AlertFlow{
				Client: flow.Client{
					Name: "test1",
					Port: 33566,
					IP:   "192.168.4.14",
				},

				Server: flow.Server{
					IP:   "104.15.15.60",
					Port: 443,
					Name: "test2",
				},
			},
			AlertProtocol: domains.AlertProtocol{
				Protocol: flow.Protocol{
					L4: "TCP",
					L7: "TLS.Google",
				},
			},
		},
	}

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix())
	ip := "192.168.4.14"

	mockService := mocks.NewMockAlertService(ctrl)
	mockHostFilter := mocks_host.NewMockHostsFilter(ctrl)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end, ip).Return(expected, nil)

	alertSearcher := usecase.NewAlertSearcher(mockService, mockHostFilter)
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
	ip := "192.168.21.1"

	mockService := mocks.NewMockAlertService(ctrl)
	mockHostFilter := mocks_host.NewMockHostsFilter(ctrl)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end, ip).Return([]domains.Alert{}, fmt.Errorf("test error"))

	alertSearcher := usecase.NewAlertSearcher(mockService, mockHostFilter)
	_, err := alertSearcher.GetAllAlerts()

	if err == nil {
		t.Fail()
	}
}
