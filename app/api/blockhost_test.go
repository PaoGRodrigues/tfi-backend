package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var Host = domains.Host{
	Name: "test.google.com",
	IP:   "192.192.192.10",
}

type blockHostRequest struct {
	Host string // Host can be IP or URL
}

func TestBlockHostByIPReturn200(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlocker := mocks.NewMockHostBlocker(ctrl)
	mockBlocker.EXPECT().Block(Host.IP).Return(Host, nil)

	api := &api.Api{
		HostBlocker: mockBlocker,
		Engine:      gin.Default(),
	}

	api.MapBlockHostURL()

	response := httptest.NewRecorder()

	req := blockHostRequest{
		Host: Host.IP,
	}

	body, _ := json.Marshal(req)

	requestUrl := "/blockhost"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBlockHostURLReturn200(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlocker := mocks.NewMockHostBlocker(ctrl)
	mockBlocker.EXPECT().Block(Host.Name).Return(Host, nil)

	api := &api.Api{
		HostBlocker: mockBlocker,
		Engine:      gin.Default(),
	}

	api.MapBlockHostURL()

	response := httptest.NewRecorder()

	req := blockHostRequest{
		Host: Host.Name,
	}

	body, _ := json.Marshal(req)

	requestUrl := "/blockhost"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBlockHostRouteReceiveWrongBodyReturn400(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	api := &api.Api{
		Engine: gin.Default(),
	}

	api.MapBlockHostURL()

	response := httptest.NewRecorder()

	req := blockHostRequest{
		Host: "",
	}

	body, _ := json.Marshal(req)

	requestUrl := "/blockhost"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestBlockHostFunctionReturningErrorReturn400(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlocker := mocks.NewMockHostBlocker(ctrl)
	mockBlocker.EXPECT().Block(Host.Name).Return(domains.Host{}, fmt.Errorf("Test error"))

	api := &api.Api{
		HostBlocker: mockBlocker,
		Engine:      gin.Default(),
	}

	api.MapBlockHostURL()

	response := httptest.NewRecorder()

	req := blockHostRequest{
		Host: Host.Name,
	}

	body, _ := json.Marshal(req)

	requestUrl := "/blockhost"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestBlockHostFunctionReturnErrorWhenTheBodyIsWrong(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlocker := mocks.NewMockHostBlocker(ctrl)

	api := &api.Api{
		HostBlocker: mockBlocker,
		Engine:      gin.Default(),
	}

	api.MapBlockHostURL()

	response := httptest.NewRecorder()

	body, _ := json.Marshal("{\"Ip\": \"10.10.10.10\"}")

	requestUrl := "/blockhost"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestBlockHostFunctionReturningErrorReturn400WhenIPNotExist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlocker := mocks.NewMockHostBlocker(ctrl)
	mockBlocker.EXPECT().Block(Host.IP).Return(domains.Host{}, nil)

	api := &api.Api{
		HostBlocker: mockBlocker,
		Engine:      gin.Default(),
	}

	api.MapBlockHostURL()

	response := httptest.NewRecorder()

	req := blockHostRequest{
		Host: Host.IP,
	}

	body, _ := json.Marshal(req)

	requestUrl := "/blockhost"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}
