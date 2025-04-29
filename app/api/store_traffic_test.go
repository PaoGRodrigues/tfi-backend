package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
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
}
var protocols = traffic.Protocol{
	L4: "UDP.Youtube",
	L7: "TLS.GoogleServices",
}

var expected = []traffic.ActiveFlow{
	traffic.ActiveFlow{
		Client:   client,
		Server:   server,
		Bytes:    1000,
		Protocol: protocols,
	},
}

func TestStoreTrafficSuccessfullyReturn200(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActiveFlowsStorage := mocks.NewMockTrafficStorage(ctrl)
	mockActiveFlowsStorage.EXPECT().StoreFlows().Return(nil)

	api := &api.Api{
		ActiveFlowsStorage: mockActiveFlowsStorage,
		Engine:             gin.Default(),
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

	mockActiveFlowsStorage := mocks.NewMockTrafficStorage(ctrl)
	mockActiveFlowsStorage.EXPECT().StoreFlows().Return(fmt.Errorf("Testing error case"))

	api := &api.Api{
		ActiveFlowsStorage: mockActiveFlowsStorage,
		Engine:             gin.Default(),
	}

	api.MapStoreActiveFlowsURL()

	response := httptest.NewRecorder()

	requestUrl := "/activeflows"
	httpRequest, _ := http.NewRequest("POST", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
