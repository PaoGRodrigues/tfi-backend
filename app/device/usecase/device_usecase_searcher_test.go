package usecase_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/device/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/device"
	"github.com/golang/mock/gomock"
)

func TestGetAllDevicesSearcherReturnAListOfDevices(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.Device{
		domains.Device{
			Name:     "Test",
			IP:       "13.13.13.13",
			OsDetail: "details",
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

func TestGetAllDevicesSearcherReturnAnError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeviceRepository := mocks.NewMockDeviceRepository(ctrl)
	mockDeviceRepository.EXPECT().GetAll().Return(nil, fmt.Errorf("Testing Error"))

	deviceSearcher := usecase.NewDeviceSearcher(mockDeviceRepository)
	_, err := deviceSearcher.GetAllDevices()

	if err == nil {
		t.Errorf("We expected an error, but didn't get one.")
	}
}
