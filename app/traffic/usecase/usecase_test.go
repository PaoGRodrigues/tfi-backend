package usecase_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/golang/mock/gomock"
)

func TestGetAllTrafficReturnAListOfTrafficJsons(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.Traffic{
		domains.Traffic{
			ID:          1234,
			Datetime:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			Source:      "192.168.4.9",
			Destination: "lib.gen.rus",
			Port:        "443",
			Protocol:    "tcp",
			Service:     "SSL",
			Bytes:       345,
		},
	}

	mockTrafficRepo := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepo.EXPECT().GetAll().Return(expected, nil)

	trafficSearcher := usecase.NewTrafficSearcher(mockTrafficRepo)
	got, err := trafficSearcher.GetAllTraffic()

	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetAllTrafficReturnAnError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrafficRepo := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepo.EXPECT().GetAll().Return(nil, fmt.Errorf("Testing Error"))

	trafficSearcher := usecase.NewTrafficSearcher(mockTrafficRepo)
	_, err := trafficSearcher.GetAllTraffic()

	if err == nil {
		t.Errorf("We expected an error, but didn't get one.")
	}
}
