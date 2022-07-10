package api_test

import (
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

var client = domains.Client{
	Name: "test",
	Port: 55672,
	IP:   "192.168.4.9",
}
var server = domains.Server{
	IP:                "123.123.123.123",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "lib.gen.rus",
}
var protocols = domains.Protocol{
	L4: "UDP.Youtube",
	L7: "TLS.GoogleServices",
}

var expected = []domains.ActiveFlow{
	domains.ActiveFlow{
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

	mockActiveFlowsStorage := mocks.NewMockActiveFlowsStorage(ctrl)
	mockActiveFlowsStorage.EXPECT().StoreFlows().Return(nil)

	api := &api.Api{
		ActiveFlowsStorage: mockActiveFlowsStorage,
		Engine:             gin.Default(),
	}

	api.MapStoreActiveFlows()

	response := httptest.NewRecorder()

	requestUrl := "/activeflows"
	httpRequest, _ := http.NewRequest("POST", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)

}
