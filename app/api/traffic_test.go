package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTrafficUseCaseAndGetAllTraffic(t *testing.T) {

	client := domains.Client{
		Name: "test",
		Port: 55672,
		IP:   "192.168.4.9",
	}
	server := domains.Server{
		IP:                "123.123.123.123",
		IsBroadcastDomain: false,
		IsDHCP:            false,
		Port:              443,
		Name:              "lib.gen.rus",
	}
	protocols := domains.Protocol{
		L4: "UDP.Youtube",
		L7: "TLS.GoogleServices",
	}

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createdTraffic := []domains.ActiveFlow{
		domains.ActiveFlow{
			Client:   client,
			Server:   server,
			Bytes:    345,
			Protocol: protocols,
		},
	}

	mockTrafficSearcherUseCase := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficSearcherUseCase.EXPECT().GetAllActiveTraffic().Return(createdTraffic, nil)

	api := &api.Api{
		TrafficSearcher: mockTrafficSearcherUseCase,
		Engine:          gin.Default(),
	}

	api.MapGetTrafficURL()

	response := httptest.NewRecorder()

	requestUrl := "/traffic"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestCreateATrafficUsecaseAndGetTrafficReturnAnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrafficSearcherUseCase := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficSearcherUseCase.EXPECT().GetAllActiveTraffic().Return(nil, fmt.Errorf("Testing error case"))

	api := &api.Api{
		TrafficSearcher: mockTrafficSearcherUseCase,
		Engine:          gin.Default(),
	}

	api.MapGetTrafficURL()

	response := httptest.NewRecorder()

	requestUrl := "/traffic"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)

}

func TestCreateTrafficActiveFlowsAndGetBytesPerDest(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		domains.BytesPerDestination{
			Bytes:       3454567,
			Destination: "google.com.ar",
		},
	}

	mockActiveFlowsSearcher := mocks.NewMockTrafficBytesParser(ctrl)
	mockActiveFlowsSearcher.EXPECT().GetBytesPerDestination().Return(expected, nil)

	api := &api.Api{
		TrafficBytesParser: mockActiveFlowsSearcher,
		Engine:             gin.Default(),
	}

	api.MapGetActiveFlowsPerDestinationURL()

	response := httptest.NewRecorder()

	requestUrl := "/activeflowsperdest"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))
	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestCreateTrafficActiveFlowsPerDestAndGetAnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActiveFlowsSearcher := mocks.NewMockTrafficBytesParser(ctrl)
	mockActiveFlowsSearcher.EXPECT().GetBytesPerDestination().Return(nil, fmt.Errorf("Testing error case"))

	api := &api.Api{
		TrafficBytesParser: mockActiveFlowsSearcher,
		Engine:             gin.Default(),
	}

	api.MapGetActiveFlowsPerDestinationURL()

	response := httptest.NewRecorder()

	requestUrl := "/activeflowsperdest"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestGetBytesPerCountryAndReturn200(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerCountry{
		domains.BytesPerCountry{
			Bytes:   3454567,
			Country: "US",
		},
	}

	mockActiveFlowsSearcher := mocks.NewMockTrafficBytesParser(ctrl)
	mockActiveFlowsSearcher.EXPECT().GetBytesPerCountry().Return(expected, nil)

	api := &api.Api{
		TrafficBytesParser: mockActiveFlowsSearcher,
		Engine:             gin.Default(),
	}

	api.MapGetActiveFlowsPerCountryURL()

	response := httptest.NewRecorder()

	requestUrl := "/activeflowspercountry"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))
	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetBytesPerCountryReturnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActiveFlowsSearcher := mocks.NewMockTrafficBytesParser(ctrl)
	mockActiveFlowsSearcher.EXPECT().GetBytesPerCountry().Return(nil, fmt.Errorf("Testing error case"))

	api := &api.Api{
		TrafficBytesParser: mockActiveFlowsSearcher,
		Engine:             gin.Default(),
	}

	api.MapGetActiveFlowsPerCountryURL()

	response := httptest.NewRecorder()

	requestUrl := "/activeflowspercountry"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))
	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
