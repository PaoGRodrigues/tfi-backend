package usecase_test

import (
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
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

	mockHostsRepo := mocks.NewMockHostsRepository(ctrl)
	mockHostsSearcher := mocks.NewMockHostUseCase(ctrl)
	mockHostsSearcher.EXPECT().GetAllHosts().Return([]domains.Host{host1, host2}, nil)
	mockHostsRepo.EXPECT().StoreHosts([]domains.Host{host1, host2}).Return(nil)

	hostsStorage := usecase.NewHostsStorage(mockHostsSearcher, mockHostsRepo)
	err := hostsStorage.StoreHosts()

	if err != nil {
		t.Fail()
	}
}
