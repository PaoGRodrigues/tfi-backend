package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	trafficUseCases "github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"
	trafficPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/traffic"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetBytesPerCountryAndReturn200(t *testing.T) {

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

	var expected = []trafficUseCases.BytesPerCountry{{
		Bytes:   5566778,
		Country: "US",
	}}

	mockTrafficDBRepository := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockTrafficDBRepository.EXPECT().GetServers().Return([]traffic.Server{server}, nil)
	mockTrafficDBRepository.EXPECT().GetFlowByKey("12344567").Return(expectedFlow, nil)

	getTrafficFlowsPerCountryUseCase := trafficUseCases.NewGetTrafficFlowsPerCountryUseCase(mockTrafficDBRepository)

	api := &api.Api{
		GetTrafficFlowsPerCountryUseCase: getTrafficFlowsPerCountryUseCase,
		Engine:                           gin.Default(),
	}

	api.MapGetActiveFlowsPerCountryURL()

	response := httptest.NewRecorder()

	requestUrl := "/activeflowspercountry"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))
	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)

	var expectedResponse = make(map[string][]trafficUseCases.BytesPerCountry)
	expectedResponse["data"] = expected

	var gotMap = make(map[string][]trafficUseCases.BytesPerCountry)
	err := json.Unmarshal(response.Body.Bytes(), &gotMap)
	require.NoError(t, err)

	assert.Equal(t, expectedResponse, gotMap)
}

func TestGetBytesPerCountryReturnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrafficDBRepository := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockTrafficDBRepository.EXPECT().GetServers().Return(nil, fmt.Errorf("Testing error case"))

	getTrafficFlowsPerCountryUseCase := trafficUseCases.NewGetTrafficFlowsPerCountryUseCase(mockTrafficDBRepository)

	api := &api.Api{
		GetTrafficFlowsPerCountryUseCase: getTrafficFlowsPerCountryUseCase,
		Engine:                           gin.Default(),
	}

	api.MapGetActiveFlowsPerCountryURL()

	response := httptest.NewRecorder()

	requestUrl := "/activeflowspercountry"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))
	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
