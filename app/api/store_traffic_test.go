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

	//	_ := mocks.NewMockTrafficUseCase(ctrl)
	mockActiveFlowsStorage := mocks.NewMockActiveFlowsStorage(ctrl)
	//	_ := mocks.NewMockTrafficRepository(ctrl)

	api := &api.Api{
		ActiveFlowsStorage: mockActiveFlowsStorage,
		Engine:             gin.Default(),
	}

	r := gin.Default()

	r.GET("/activeflows", api.StoreActiveTraffic,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/activeflows"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	t.Run("Ok", func(t *testing.T) {
		mockActiveFlowsStorage.EXPECT().StoreFlows().Return(nil)

		res := executeWithContext()
		assert.Equal(t, http.StatusOK, res.Code)
	})
}
