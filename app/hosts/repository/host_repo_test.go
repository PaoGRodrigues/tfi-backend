package repository_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/repository"
	services_mocks "github.com/PaoGRodrigues/tfi-backend/mocks/services"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

var host1 = host.Host{
	Name:        "test",
	PrivateHost: false,
	IP:          "123.123.123.123",
	City:        "",
	Country:     "US",
}

var host2 = host.Host{
	Name:        "test.randomdns.com",
	PrivateHost: false,
	IP:          "13.13.13.13",
	City:        "BuenosAires",
	Country:     "AR",
}

func TestStoreHostsSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().AddHosts([]host.Host{host1, host2}).Return(nil)

	hostsStorage := repository.NewHostsRepo(mockDatabase)
	err := hostsStorage.StoreHosts([]host.Host{host1, host2})

	if err != nil {
		t.Fail()
	}
}

func TestStoreHostsWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().AddHosts([]host.Host{host1, host2}).Return(fmt.Errorf("Test error"))

	hostsStorage := repository.NewHostsRepo(mockDatabase)
	err := hostsStorage.StoreHosts([]host.Host{host1, host2})

	if err == nil {
		t.Fail()
	}
}

func TestGetHostReturnHostSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetHostByIp(host1.IP).Return(host1, nil)

	hostStorage := repository.NewHostsRepo(mockDatabase)
	got, err := hostStorage.GetHostByIp(host1.IP)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, host1, got)
}

func TestGetHostReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetHostByIp(host1.IP).Return(host.Host{}, fmt.Errorf("test Error"))

	hostStorage := repository.NewHostsRepo(mockDatabase)
	_, err := hostStorage.GetHostByIp(host1.IP)

	if err == nil {
		t.Fail()
	}
}
