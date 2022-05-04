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

	mockTrafficSearcherUseCase := mocks.NewMockTrafficUseCase(ctrl)

	api := &api.Api{
		TrafficSearcher: mockTrafficSearcherUseCase,
		Engine:          gin.Default(),
	}

	r := gin.Default()

	r.GET("/traffic", api.GetTraffic,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/traffic"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	createdTraffic := []domains.ActiveFlow{
		domains.ActiveFlow{
			Client:   client,
			Server:   server,
			Bytes:    345,
			Protocol: protocols,
		},
	}

	t.Run("Ok", func(t *testing.T) {

		mockTrafficSearcherUseCase.EXPECT().GetAllActiveTraffic().Return(createdTraffic, nil)

		res := executeWithContext()
		assert.Equal(t, http.StatusOK, res.Code)
	})
}

func TestCreateATrafficUsecaseAndGetTrafficReturnAnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrafficSearcherUseCase := mocks.NewMockTrafficUseCase(ctrl)

	api := &api.Api{
		TrafficSearcher: mockTrafficSearcherUseCase,
		Engine:          gin.Default(),
	}

	r := gin.Default()

	r.GET("/traffic", api.GetTraffic,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/traffic"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	t.Run("Ok", func(t *testing.T) {

		mockTrafficSearcherUseCase.EXPECT().GetAllActiveTraffic().Return(nil, fmt.Errorf("Testing error case"))

		res := executeWithContext()
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestCreateTrafficActiveFlowsAndGetBytesPerDest(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActiveFlowsSearcher := mocks.NewMockTrafficActiveFlowsSearcher(ctrl)

	api := &api.Api{
		ActiveFlowsSearcher: mockActiveFlowsSearcher,
		Engine:              gin.Default(),
	}

	r := gin.Default()

	r.GET("/activeflowsperdest", api.GetActiveFlowsPerDestination,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/activeflowsperdest"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	expected := []domains.BytesPerDestination{
		domains.BytesPerDestination{
			Bytes:       3454567,
			Destination: "google.com.ar",
		},
	}

	t.Run("Ok", func(t *testing.T) {

		mockActiveFlowsSearcher.EXPECT().GetBytesPerDestination().Return(expected, nil)

		res := executeWithContext()
		assert.Equal(t, http.StatusOK, res.Code)
	})
}

func TestCreateTrafficActiveFlowsAndGetAnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActiveFlowsSearcher := mocks.NewMockTrafficActiveFlowsSearcher(ctrl)

	api := &api.Api{
		ActiveFlowsSearcher: mockActiveFlowsSearcher,
		Engine:              gin.Default(),
	}

	r := gin.Default()

	r.GET("/activeflowsperdest", api.GetActiveFlowsPerDestination,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/activeflowsperdest"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	t.Run("Ok", func(t *testing.T) {

		mockActiveFlowsSearcher.EXPECT().GetBytesPerDestination().Return(nil, fmt.Errorf("Testing error case"))

		res := executeWithContext()
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
