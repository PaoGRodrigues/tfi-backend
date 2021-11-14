package device_test

import (
	"reflect"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/device/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/tests/mocks/device"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestGetAllDevicesSearcherReturnAListOfDevices(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.Device{
		domains.Device{
			ID:      1,
			Name:    "Test",
			IP:      "13.13.13.13",
			Details: "details",
		},
	}

	mockDeviceRepository := mocks.NewMockDeviceRepository(ctrl)
	mockDeviceRepository.EXPECT().GetAll().Return(expected, nil)

	deviceSearcher := usecase.NewDeviceSearcher(mockDeviceRepository)
	got, err := deviceSearcher.GetAllDevices()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}
