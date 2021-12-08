package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTrafficUseCaseAndGetAllTraffic(t *testing.T) {

	var (
		id          = 1
		datetime    = time.Now()
		source      = "172.16.0.0"
		destination = "8.8.8.8"
		port        = 443
		protocol    = "tcp"
		service     = "Ssl"
		bytes       = 234567
	)

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrafficSearcherUseCase := mocks.NewMockTrafficUseCase(ctrl)

	api := &api.Api{
		TrafficUseCase: mockTrafficSearcherUseCase,
		Engine:         gin.Default(),
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

	createdTraffic := []domains.Traffic{
		domains.Traffic{
			ID:          id,
			Datetime:    datetime,
			Source:      source,
			Destination: destination,
			Port:        port,
			Protocol:    protocol,
			Service:     service,
			Bytes:       bytes,
		},
	}

	t.Run("Ok", func(t *testing.T) {

		mockTrafficSearcherUseCase.EXPECT().GetAllTraffic().Return(createdTraffic, nil)

		res := executeWithContext()
		assert.Equal(t, http.StatusOK, res.Code)
	})
}
