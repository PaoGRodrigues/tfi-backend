package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	flow "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	alertUseCase "github.com/PaoGRodrigues/tfi-backend/app/usecase/alert"
	alertPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/alert"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateAlertsUsecaseGetAllAlertsReturnAlerts(t *testing.T) {

	expected := []alert.Alert{
		alert.Alert{

			Name:     "test",
			Family:   "flow",
			Time:     "10/10/10 11:11:11",
			Severity: "Advertencia",
			AlertFlow: alert.AlertFlow{
				Client: flow.Client{
					Port: 33566,
					IP:   "192.168.4.14",
					Name: "192.168.4.14",
				},

				Server: flow.Server{
					IP:   "104.15.15.60",
					Port: 443,
					Name: "test2",
				},
			},
			AlertProtocol: flow.Protocol{
				L4:    "TCP",
				L7:    "TLS.Google",
				Label: "TCP:TLS.Google",
			},
		},
	}

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := alertPortsMock.NewMockAlertReader(ctrl)
	mockRepository.EXPECT().GetAllAlerts(gomock.Any(), gomock.Any()).Return(expected, nil)

	getAlertsUseCase := alertUseCase.NewGetAlertsUseCase(mockRepository)

	api := &api.Api{
		GetAlertsUseCase: getAlertsUseCase,
		Engine:           gin.Default(),
	}

	api.MapAlertsURL()

	response := httptest.NewRecorder()

	requestUrl := "/alerts"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCreateAlertsUsecaseGetAllAlertsReturnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := alertPortsMock.NewMockAlertReader(ctrl)
	mockRepository.EXPECT().GetAllAlerts(gomock.Any(), gomock.Any()).Return([]alert.Alert{}, fmt.Errorf("Error test"))

	getAlertsUseCase := alertUseCase.NewGetAlertsUseCase(mockRepository)

	api := &api.Api{
		GetAlertsUseCase: getAlertsUseCase,
		Engine:           gin.Default(),
	}

	api.MapAlertsURL()

	response := httptest.NewRecorder()

	requestUrl := "/alerts"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestCreateAlertsUsecaseGetAllAlertsWithDestNameEmptyReturnAlerts(t *testing.T) {

	expected := []alert.Alert{
		alert.Alert{

			Name:     "test",
			Family:   "flow",
			Time:     "10/10/10 11:11:11",
			Severity: "Advertencia",
			AlertFlow: alert.AlertFlow{
				Client: flow.Client{
					Port: 33566,
					IP:   "192.168.4.14",
					Name: "192.168.4.14",
				},

				Server: flow.Server{
					IP:   "104.15.15.60",
					Port: 443,
					Name: "",
				},
			},
			AlertProtocol: flow.Protocol{
				L4:    "TCP",
				L7:    "TLS.Google",
				Label: "TCP:TLS.Google",
			},
		},
	}

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := alertPortsMock.NewMockAlertReader(ctrl)
	mockRepository.EXPECT().GetAllAlerts(gomock.Any(), gomock.Any()).Return(expected, nil)

	getAlertsUseCase := alertUseCase.NewGetAlertsUseCase(mockRepository)

	api := &api.Api{
		GetAlertsUseCase: getAlertsUseCase,
		Engine:           gin.Default(),
	}

	api.MapAlertsURL()

	response := httptest.NewRecorder()

	requestUrl := "/alerts"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}
