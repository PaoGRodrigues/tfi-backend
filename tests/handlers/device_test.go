package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/handlers"
	"github.com/PaoGRodrigues/tfi-backend/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateADeviceHandlerAndGetAllDevicesFromAUseCase(t *testing.T) {

	var (
		id      = 1
		name    = "Test"
		ip      = "13.13.13.13"
		details = "details"
	)

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeviceUseCase := mocks.NewMockDeviceUseCase(ctrl)

	executeWithContext := func(MockDeviceUseCase *mocks.MockDeviceUseCase) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)

		requestUrl := "/devices"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		handlers.NewDeviceHandler(ginEngine, MockDeviceUseCase)
		ginEngine.ServeHTTP(response, httpRequest)
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

		mockDeviceUseCase.EXPECT().GetAll(gomock.Any()).Return(createdDevices, nil)

		res := executeWithContext(mockDeviceUseCase)
		assert.Equal(t, http.StatusOK, res.Code)
	})
}
