package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAlertsUsecaseGetAllAlertsReturnAlerts(t *testing.T) {

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
