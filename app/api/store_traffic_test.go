package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	trafficUseCase "github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"
	trafficPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/traffic"

	hostPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/host"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var client = traffic.Client{
	Name: "test",
	Port: 55672,
	IP:   "192.168.4.9",
}
var server = traffic.Server{
	IP:                "123.123.123.123",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "lib.gen.rus",
	Country:           "US",
}
var protocols = traffic.Protocol{
	L4: "UDP.Youtube",
	L7: "TLS.GoogleServices",
}

var expected = []traffic.TrafficFlow{
	traffic.TrafficFlow{
		Client:   client,
		Server:   server,
		Bytes:    1000,
		Protocol: protocols,
	},
}

var hostExpected = host.Host{
	Name:        "test",
	PrivateHost: false,
	IP:          "123.123.123.123",
	City:        "",
	Country:     "US",
}

func TestStoreTrafficSuccessfullyReturn200(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrafficReader := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockTrafficReader.EXPECT().GetTrafficFlows().Return(expected, nil)

	mockHostReader := hostPortsMock.NewMockHostDBRepository(ctrl)
	mockHostReader.EXPECT().GetHostByIp(expected[0].Server.IP).Return(hostExpected, nil)

	mockStoreTrafficFlows := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockStoreTrafficFlows.EXPECT().StoreTrafficFlows(expected).Return(nil)

	trafficUseCase := trafficUseCase.NewStoreTrafficFlowsUseCase(mockTrafficReader, mockStoreTrafficFlows, mockHostReader)

	api := &api.Api{
		StoreTrafficFlowsUseCase: trafficUseCase,
		Engine:                   gin.Default(),
	}

	api.MapStoreActiveFlowsURL()

	response := httptest.NewRecorder()

	requestUrl := "/activeflows"
	httpRequest, _ := http.NewRequest("POST", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestStoreTrafficFailedAndReturn500(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrafficReader := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockTrafficReader.EXPECT().GetTrafficFlows().Return(expected, nil)

	mockStoreTrafficFlows := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockStoreTrafficFlows.EXPECT().StoreTrafficFlows(expected).Return(fmt.Errorf("Testing error case"))

	mockHostReader := hostPortsMock.NewMockHostDBRepository(ctrl)
	mockHostReader.EXPECT().GetHostByIp(expected[0].Server.IP).Return(hostExpected, nil)

	trafficUseCase := trafficUseCase.NewStoreTrafficFlowsUseCase(mockTrafficReader, mockStoreTrafficFlows, mockHostReader)

	api := &api.Api{
		StoreTrafficFlowsUseCase: trafficUseCase,
		Engine:                   gin.Default(),
	}

	api.MapStoreActiveFlowsURL()

	response := httptest.NewRecorder()

	requestUrl := "/activeflows"
	httpRequest, _ := http.NewRequest("POST", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
