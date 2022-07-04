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

	mockTrafficSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficSearcher.EXPECT().GetAllActiveTraffic().Return(expected, nil)

	api := &api.Api{
		TrafficSearcher: mockTrafficSearcher,
		Engine:          gin.Default(),
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

		mockTrafficSearcher.EXPECT().GetAllActiveTraffic().Return(nil, fmt.Errorf("Testing error case"))

		res := executeWithContext()
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

}
