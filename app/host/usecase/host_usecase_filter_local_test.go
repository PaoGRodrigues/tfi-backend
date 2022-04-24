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

func TestGetLocalHostWithHostsReturnedFromSearcherReturnLocalHosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	local := domains.Host{
		Name:        "Test",
		IP:          "13.13.13.13",
		PrivateHost: true,
	}
	remote := domains.Host{
		Name:        "Test2",
		IP:          "172.172.172.172",
		PrivateHost: false,
	}

	expected := []domains.Host{
		local,
		remote,
	}

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return(expected)

	filter := usecase.NewLocalHosts(mockSearcher)
	got, err := filter.GetLocalHosts()
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual([]domains.Host{local}, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", local, got)
	}
}

func TestGetLocalHostCallingGetHostFromRepoInSearcherReturnLocalHosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	local := domains.Host{
		Name:        "Test",
		IP:          "13.13.13.13",
		PrivateHost: true,
	}
	remote := domains.Host{
		Name:        "Test2",
		IP:          "172.172.172.172",
		PrivateHost: false,
	}

	expected := []domains.Host{
		local,
		remote,
	}

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return([]domains.Host{})
	mockSearcher.EXPECT().GetAllHosts().Return(expected, nil)

	filter := usecase.NewLocalHosts(mockSearcher)
	got, err := filter.GetLocalHosts()
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual([]domains.Host{local}, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", local, got)
	}
}

func TestGetLocalHostAndGetAllHostsInSearcherReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return([]domains.Host{})
	mockSearcher.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing Error"))

	filter := usecase.NewLocalHosts(mockSearcher)
	_, err := filter.GetLocalHosts()

	if err == nil {
		t.Fail()
	}
}
