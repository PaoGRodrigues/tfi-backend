package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/alerts"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestSendMessageReturn200(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSender := mocks.NewMockAlertsSender(ctrl)
	mockSender.EXPECT().SendLastAlertMessages().Return(nil)

	api := &api.Api{
		AlertsSender: mockSender,
		Engine:       gin.Default(),
	}

	api.MapNotificationsURL()

	response := httptest.NewRecorder()

	requestURL := "/alertnotification"
	httpRequest, _ := http.NewRequest("GET", requestURL, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSendMessageReturn500Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSender := mocks.NewMockAlertsSender(ctrl)
	mockSender.EXPECT().SendLastAlertMessages().Return(fmt.Errorf("Error test"))

	api := &api.Api{
		AlertsSender: mockSender,
		Engine:       gin.Default(),
	}

	api.MapNotificationsURL()

	response := httptest.NewRecorder()

	requestURL := "/alertnotification"
	httpRequest, _ := http.NewRequest("GET", requestURL, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
