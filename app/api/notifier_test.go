package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	alertUseCase "github.com/PaoGRodrigues/tfi-backend/app/usecase/alert"
	alertPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/alert"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func TestSendMessageReturn200(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNotifier := alertPortsMock.NewMockNotifier(ctrl)
	mockNotifier.EXPECT().SendMessage("Alerta 1").Return(nil).AnyTimes()

	mockRepository := alertPortsMock.NewMockAlertReader(ctrl)
	mockRepository.EXPECT().GetAllAlerts(gomock.Any(), gomock.Any()).Return([]alert.Alert{{Name: "Alerta 1"}}, nil)

	notifyAlertsUseCase := alertUseCase.NewNotifyAlertsUseCase(mockNotifier, mockRepository)

	api := &api.Api{
		NotifyAlertsUseCase: notifyAlertsUseCase,
		Engine:              gin.Default(),
	}

	api.MapNotificationsURL()

	response := httptest.NewRecorder()

	requestURL := "/alertnotification"
	httpRequest, _ := http.NewRequest("POST", requestURL, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSendMessageReturn500Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := alertPortsMock.NewMockAlertReader(ctrl)
	mockRepository.EXPECT().GetAllAlerts(gomock.Any(), gomock.Any()).Return([]alert.Alert{{Name: "Alerta 2"}}, nil)
	mockNotifier := alertPortsMock.NewMockNotifier(ctrl)
	mockNotifier.EXPECT().SendMessage("Alerta 2").Return()

	notifyAlertsUseCase := alertUseCase.NewNotifyAlertsUseCase(mockNotifier, mockRepository)

	api := &api.Api{
		NotifyAlertsUseCase: notifyAlertsUseCase,
		Engine:              gin.Default(),
	}

	api.MapNotificationsURL()

	response := httptest.NewRecorder()

	requestURL := "/alertnotification"
	httpRequest, _ := http.NewRequest("POST", requestURL, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
