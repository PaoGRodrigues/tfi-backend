package usecase_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	"github.com/golang/mock/gomock"
)

func TestGetAllHostsSearcherReturnAListOfHosts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []host.Host{
		host.Host{
			Name: "Test",
			IP:   "13.13.13.13",
		},
	}

	mockHostService := mocks.NewMockHostService(ctrl)
	mockHostService.EXPECT().GetAllHosts().Return(expected, nil)

	HostSearcher := usecase.NewHostSearcher(mockHostService)
	got, err := HostSearcher.GetAllHosts()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetAllHostsSearcherReturnAnError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostService := mocks.NewMockHostService(ctrl)
	mockHostService.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing Error"))

	HostSearcher := usecase.NewHostSearcher(mockHostService)
	_, err := HostSearcher.GetAllHosts()

	if err == nil {
		t.Errorf("We expected an error, but didn't get one.")
	}
}
