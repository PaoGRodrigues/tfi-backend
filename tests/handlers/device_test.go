package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/device/gateway"
	"github.com/PaoGRodrigues/tfi-backend/app/device/handlers"

	mocks "github.com/PaoGRodrigues/tfi-backend/tests/mocks/device"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateADeviceHandlerAndGetAllDevicesFromTheGateway(t *testing.T) {

	var (
		id      = 1
		name    = "Test"
		ip      = "13.13.13.13"
		details = "details"
	)

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeviceRepository := mocks.NewMockDeviceRepository(ctrl)
	mockDeviceGateway := gateway.NewDeviceSearcher(mockDeviceRepository)
	mockDeviceHandler := handlers.NewDeviceHandler(mockDeviceGateway)

	api := &api.Api{
		DeviceHandler: mockDeviceHandler,
		Engine:        gin.Default(),
	}

	r := gin.Default()

	r.GET("/devices", api.GetDevices,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/devices"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	createdDevices := []domains.Device{
		domains.Device{
			ID:      id,
			Name:    name,
			IP:      ip,
			Details: details,
		},
	}

	t.Run("Ok", func(t *testing.T) {

		mockDeviceRepository.EXPECT().GetAll().Return(createdDevices, nil)

		res := executeWithContext()
		assert.Equal(t, http.StatusOK, res.Code)
	})
}
