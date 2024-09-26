package repository_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/repository"
	services_mocks "github.com/PaoGRodrigues/tfi-backend/mocks/services"
	"github.com/golang/mock/gomock"
)

var host1 = domains.Host{
	Name:        "test",
	PrivateHost: false,
	IP:          "123.123.123.123",
	City:        "",
	Country:     "US",
}

var host2 = domains.Host{
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
	mockDatabase.EXPECT().AddHosts([]domains.Host{host1, host2}).Return(nil)

	hostsStorage := repository.NewHostsRepo(mockDatabase)
	err := hostsStorage.StoreHosts([]domains.Host{host1, host2})

	if err != nil {
		t.Fail()
	}
}

func TestStoreHostsWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().AddHosts([]domains.Host{host1, host2}).Return(fmt.Errorf("Test error"))

	hostsStorage := repository.NewHostsRepo(mockDatabase)
	err := hostsStorage.StoreHosts([]domains.Host{host1, host2})

	if err == nil {
		t.Fail()
	}
}
