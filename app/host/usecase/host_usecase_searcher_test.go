package usecase_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/host/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/host/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/host"
	"github.com/golang/mock/gomock"
)

func TestGetAllHostsSearcherReturnAListOfHosts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.Host{
		domains.Host{
			Name: "Test",
			IP:   "13.13.13.13",
		},
	}

	mockHostRepository := mocks.NewMockHostRepository(ctrl)
	mockHostRepository.EXPECT().GetAllHosts().Return(expected, nil)

	HostSearcher := usecase.NewHostSearcher(mockHostRepository)
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

	mockHostRepository := mocks.NewMockHostRepository(ctrl)
	mockHostRepository.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing Error"))

	HostSearcher := usecase.NewHostSearcher(mockHostRepository)
	_, err := HostSearcher.GetAllHosts()

	if err == nil {
		t.Errorf("We expected an error, but didn't get one.")
	}
}
