package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	trafficUseCases "github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"
	trafficPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/traffic"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTrafficUseCaseAndGetAllTraffic(t *testing.T) {

	client := traffic.Client{
		Name: "test",
		Port: 55672,
		IP:   "192.168.4.9",
	}
	server := traffic.Server{
		IP:                "123.123.123.123",
		IsBroadcastDomain: false,
		IsDHCP:            false,
		Port:              443,
		Name:              "lib.gen.rus",
	}
	protocols := traffic.Protocol{
		L4: "UDP.Youtube",
		L7: "TLS.GoogleServices",
	}

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createdTraffic := []traffic.TrafficFlow{
		traffic.TrafficFlow{
			Client:   client,
			Server:   server,
			Bytes:    345,
			Protocol: protocols,
		},
	}

	mockReader := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockReader.EXPECT().GetTrafficFlows().Return(createdTraffic, nil)

	getTrafficFlowsUseCase := trafficUseCases.NewTrafficFlowsUseCase(mockReader)

	api := &api.Api{
		TrafficSearcher: getTrafficFlowsUseCase,
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

	mockTrafficRepository := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockTrafficRepository.EXPECT().GetTrafficFlows().Return(nil, fmt.Errorf("Testing error case"))

	GetTrafficFlowsUseCase := trafficUseCases.NewTrafficFlowsUseCase(mockTrafficRepository)

	api := &api.Api{
		TrafficSearcher: GetTrafficFlowsUseCase,
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

	var server = traffic.Server{
		IP:                "8.8.8.8",
		IsBroadcastDomain: false,
		IsDHCP:            false,
		Port:              443,
		Name:              "google.com.ar",
		Country:           "US",
		Key:               "12344567",
	}

	var expectedFlow = traffic.TrafficFlow{
		Client: traffic.Client{
			Name: "Local",
			Port: 12345,
			IP:   "192.168.4.1",
		},
		Server: server,
		Protocol: traffic.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Bytes: 5566778,
	}

	mockTrafficDBRepository := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockTrafficDBRepository.EXPECT().GetServers().Return([]traffic.Server{server}, nil)
	mockTrafficDBRepository.EXPECT().GetFlowByKey("12344567").Return(expectedFlow, nil)
	getTrafficFlowsPerDestinationUseCase := trafficUseCases.NewGetTrafficFlowsPerDestinationUseCase(mockTrafficDBRepository)

	api := &api.Api{
		GetTrafficFlowsPerDestinationUseCase: getTrafficFlowsPerDestinationUseCase,
		Engine:                               gin.Default(),
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

	mockTrafficDBRepository := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockTrafficDBRepository.EXPECT().GetServers().Return(nil, fmt.Errorf("Testing error case"))
	getTrafficFlowsPerDestinationUseCase := trafficUseCases.NewGetTrafficFlowsPerDestinationUseCase(mockTrafficDBRepository)

	api := &api.Api{
		GetTrafficFlowsPerDestinationUseCase: getTrafficFlowsPerDestinationUseCase,
		Engine:                               gin.Default(),
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

	expected := []traffic.BytesPerCountry{
		traffic.BytesPerCountry{
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
