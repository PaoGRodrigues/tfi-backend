package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/device"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeviceUseCaseAndGetAllDevices(t *testing.T) {

	var (
		id      = 1
		name    = "Test"
		ip      = "13.13.13.13"
		details = "details"
	)

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeviceSearcherUseCase := mocks.NewMockDeviceUseCase(ctrl)

	api := &api.Api{
		DeviceUseCase: mockDeviceSearcherUseCase,
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

		mockDeviceSearcherUseCase.EXPECT().GetAllDevices().Return(createdDevices, nil)

		res := executeWithContext()
		assert.Equal(t, http.StatusOK, res.Code)
	})
}

func TestCreateADeviceUsecaseAndGetDevicesReturnAnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeviceSearcherUseCase := mocks.NewMockDeviceUseCase(ctrl)

	api := &api.Api{
		DeviceUseCase: mockDeviceSearcherUseCase,
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

	t.Run("Ok", func(t *testing.T) {

		mockDeviceSearcherUseCase.EXPECT().GetAllDevices().Return(nil, fmt.Errorf("Testing error case"))

		res := executeWithContext()
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
