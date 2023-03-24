package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type configRequest struct {
	Token    string `json:"token" binding:"required"`
	Username string `json:"username" binding:"required"`
}

var config = configRequest{
	Token:    "asfklaet12124443:alllaromms",
	Username: "user123",
}

func TestConfigureReturn200(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNotiChannel := mocks.NewMockNotificationChannel(ctrl)
	mockNotiChannel.EXPECT().Configure(config.Token, config.Username).Return(nil)

	api := &api.Api{
		NotifChannel: mockNotiChannel,
		Engine:       gin.Default(),
	}

	api.MapConfigureNotifChannelURL()

	response := httptest.NewRecorder()

	req := configRequest{
		Token:    config.Token,
		Username: config.Username,
	}

	body, _ := json.Marshal(req)

	requestUrl := "/configurechannel"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestConfigurePostRequestWithWrongBodyReturn400(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNotiChannel := mocks.NewMockNotificationChannel(ctrl)

	api := &api.Api{
		NotifChannel: mockNotiChannel,
		Engine:       gin.Default(),
	}

	api.MapConfigureNotifChannelURL()

	response := httptest.NewRecorder()

	req := configRequest{
		Token:    config.Token,
		Username: "",
	}

	body, _ := json.Marshal(req)

	requestUrl := "/configurechannel"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestConfigurePostRequestReturnErrorInConfigureFunctionAndReturn500(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNotiChannel := mocks.NewMockNotificationChannel(ctrl)
	mockNotiChannel.EXPECT().Configure(config.Token, config.Username).Return(fmt.Errorf("Testing error"))

	api := &api.Api{
		NotifChannel: mockNotiChannel,
		Engine:       gin.Default(),
	}

	api.MapConfigureNotifChannelURL()

	response := httptest.NewRecorder()

	req := configRequest{
		Token:    config.Token,
		Username: config.Username,
	}

	body, _ := json.Marshal(req)

	requestUrl := "/configurechannel"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
