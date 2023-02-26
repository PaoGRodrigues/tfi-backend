package usecase_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/alerts/usecase"
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	mocks_traffic "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

var expected = []domains.Alert{
	domains.Alert{
		Name:     "test",
		Family:   "flow",
		Time:     struct{ Label string }{"10/10/10 11:11:11"},
		Score:    "10",
		Severity: domains.Severity{Label: "2"},
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
	domains.Alert{
		Name:     "test",
		Family:   "flow",
		Time:     struct{ Label string }{"10/10/10 11:11:11"},
		Score:    "10",
		Severity: domains.Severity{Label: "2"},
		AlertFlow: domains.AlertFlow{
			Client: flow.Client{
				Name: "test2",
				Port: 33566,
				IP:   "192.168.4.15",
			},

			Server: flow.Server{
				IP:   "104.15.15.70",
				Port: 443,
				Name: "test3",
			},
		},
		AlertProtocol: domains.AlertProtocol{
			Protocol: flow.Protocol{
				L4: "TCP",
				L7: "TLS.YouTube",
			},
		},
	},
}

func TestGetAllAlertsReturnListOfAlertsWhenServersListIsEmpty(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix())
	ip1 := "192.168.4.14"

	mockService := mocks.NewMockAlertService(ctrl)
	mockTrafficFilter := mocks_traffic.NewMockActiveFlowsStorage(ctrl)
	mockTrafficFilter.EXPECT().GetClientsList().Return([]flow.Client{flow.Client{IP: ip1}}, nil)
	mockTrafficFilter.EXPECT().GetServersList().Return(nil, nil)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end, ip1).Return(expected, nil)

	alertSearcher := usecase.NewAlertSearcher(mockService, mockTrafficFilter)
	got, err := alertSearcher.GetAllAlerts()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected, got)
}

func TestGetAllAlertsReturnListOfAlertsWhenClientsListIsEmpty(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix())
	ip1 := "192.168.4.14"

	mockService := mocks.NewMockAlertService(ctrl)
	mockTrafficFilter := mocks_traffic.NewMockActiveFlowsStorage(ctrl)
	mockTrafficFilter.EXPECT().GetClientsList().Return(nil, nil)
	mockTrafficFilter.EXPECT().GetServersList().Return([]flow.Server{flow.Server{IP: ip1}}, nil)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end, ip1).Return(expected, nil)

	alertSearcher := usecase.NewAlertSearcher(mockService, mockTrafficFilter)
	got, err := alertSearcher.GetAllAlerts()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected, got)
}

func TestGetAllAlertsReturnErrorWhenCallServiceForClients(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix())
	ip := "192.168.21.1"

	mockService := mocks.NewMockAlertService(ctrl)
	mockTrafficFilter := mocks_traffic.NewMockActiveFlowsStorage(ctrl)
	mockTrafficFilter.EXPECT().GetClientsList().Return([]flow.Client{{IP: ip}}, nil)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end, ip).Return([]domains.Alert{}, fmt.Errorf("test error"))

	alertSearcher := usecase.NewAlertSearcher(mockService, mockTrafficFilter)
	_, err := alertSearcher.GetAllAlerts()

	if err == nil {
		t.Fail()
	}
}

func TestGetAllAlertsReturnErrorWhenCallGetAllAlertsForClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAlertService(ctrl)
	mockTrafficFilter := mocks_traffic.NewMockActiveFlowsStorage(ctrl)
	mockTrafficFilter.EXPECT().GetClientsList().Return(nil, fmt.Errorf("test error"))

	alertSearcher := usecase.NewAlertSearcher(mockService, mockTrafficFilter)
	_, err := alertSearcher.GetAllAlerts()

	if err == nil {
		t.Fail()
	}
}

func TestGetAllAlertsReturnErrorWhenCallGetAllAlertsForServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAlertService(ctrl)
	mockTrafficFilter := mocks_traffic.NewMockActiveFlowsStorage(ctrl)
	mockTrafficFilter.EXPECT().GetClientsList().Return(nil, nil)
	mockTrafficFilter.EXPECT().GetServersList().Return(nil, fmt.Errorf("test error"))

	alertSearcher := usecase.NewAlertSearcher(mockService, mockTrafficFilter)
	_, err := alertSearcher.GetAllAlerts()

	if err == nil {
		t.Fail()
	}
}

func TestGetAllAlertsReturnErrorWhenCallServiceForServers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix())
	ip := "192.168.21.1"

	mockService := mocks.NewMockAlertService(ctrl)
	mockTrafficFilter := mocks_traffic.NewMockActiveFlowsStorage(ctrl)
	mockTrafficFilter.EXPECT().GetClientsList().Return(nil, nil)
	mockTrafficFilter.EXPECT().GetServersList().Return([]flow.Server{{IP: ip}}, nil)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end, ip).Return([]domains.Alert{}, fmt.Errorf("test error"))

	alertSearcher := usecase.NewAlertSearcher(mockService, mockTrafficFilter)
	_, err := alertSearcher.GetAllAlerts()

	if err == nil {
		t.Fail()
	}
}

func TestGetAllAlertsReturnClientAndServerAlerts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	epoch_end := int(now.Unix())
	epoch_begin := int(now.AddDate(0, 0, -7).Unix())
	ip1 := "192.168.21.1"
	name2 := "test3"
	ip2 := "104.15.15.70"

	mockService := mocks.NewMockAlertService(ctrl)
	mockTrafficFilter := mocks_traffic.NewMockActiveFlowsStorage(ctrl)
	mockTrafficFilter.EXPECT().GetClientsList().Return([]flow.Client{{IP: ip1}}, nil)
	mockTrafficFilter.EXPECT().GetServersList().Return([]flow.Server{{Name: name2, IP: ip2}}, nil)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end, ip1).Return([]domains.Alert{expected[0]}, nil)
	mockService.EXPECT().GetAllAlerts(epoch_begin, epoch_end, ip2).Return([]domains.Alert{expected[1]}, nil)

	alertSearcher := usecase.NewAlertSearcher(mockService, mockTrafficFilter)
	got, err := alertSearcher.GetAllAlerts()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected, got)
}
