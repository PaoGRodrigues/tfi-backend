package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAlertsUsecaseGetAllAlertsReturnAlerts(t *testing.T) {

	expected := []domains.Alert{
		domains.Alert{

			Name:     "test",
			Family:   "flow",
			Time:     struct{ Label string }{"10/10/10 11:11:11"},
			Severity: domains.Severity{Label: "2"},
			AlertFlow: domains.AlertFlow{
				Client: domains.AlertClient{
					CliPort: 33566,
					Value:   "192.168.4.14",
				},

				Server: domains.AlertServer{
					Value:   "104.15.15.60",
					SrvPort: 443,
					Name:    "test2",
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

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlertSearcher := mocks.NewMockAlertUseCase(ctrl)
	mockAlertSearcher.EXPECT().GetAllAlerts().Return(expected, nil)

	api := &api.Api{
		AlertsSearcher: mockAlertSearcher,
		Engine:         gin.Default(),
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

	mockAlertSearcher := mocks.NewMockAlertUseCase(ctrl)
	mockAlertSearcher.EXPECT().GetAllAlerts().Return([]domains.Alert{}, fmt.Errorf("Error test"))

	api := &api.Api{
		AlertsSearcher: mockAlertSearcher,
		Engine:         gin.Default(),
	}

	api.MapAlertsURL()

	response := httptest.NewRecorder()

	requestUrl := "/alerts"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
