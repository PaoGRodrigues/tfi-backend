package usecase_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	"github.com/golang/mock/gomock"
)

var local = domains.Host{
	Name:        "Test",
	IP:          "13.13.13.13",
	PrivateHost: true,
}
var remote = domains.Host{
	Name:        "Test2",
	IP:          "172.172.172.172",
	PrivateHost: false,
}

var expected = []domains.Host{
	local,
	remote,
}

func TestGetLocalHostWithHostsReturnedFromSearcherReturnLocalHosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return(expected)

	filter := usecase.NewHostsFilter(mockSearcher)
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

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return([]domains.Host{})
	mockSearcher.EXPECT().GetAllHosts().Return(expected, nil)

	filter := usecase.NewHostsFilter(mockSearcher)
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

	filter := usecase.NewHostsFilter(mockSearcher)
	_, err := filter.GetLocalHosts()

	if err == nil {
		t.Fail()
	}
}

func TestGetRemoteHostAndCallGetAllHostsInSearcherReturnRemoteHostsSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return(expected)

	filter := usecase.NewHostsFilter(mockSearcher)
	got, err := filter.GetRemoteHosts()
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual([]domains.Host{remote}, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", remote, got)
	}

}

func TestGetRemoteHostCallingGetHostFromRepoInSearcherReturnRemoteHosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return([]domains.Host{})
	mockSearcher.EXPECT().GetAllHosts().Return(expected, nil)

	filter := usecase.NewHostsFilter(mockSearcher)
	got, err := filter.GetRemoteHosts()
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual([]domains.Host{remote}, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", remote, got)
	}
}

func TestGetRemoteHostAndGetAllHostsInSearcherReturnAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return([]domains.Host{})
	mockSearcher.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing Error"))

	filter := usecase.NewHostsFilter(mockSearcher)
	_, err := filter.GetRemoteHosts()

	if err == nil {
		t.Fail()
	}
}
