package host_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	usecase "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"

	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	hostPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/host"
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

	mockHostsRepo := mocks.NewMockHostsRepository(ctrl)

	mockRepository := hostPortsMock.NewMockHostRepository(ctrl)
	mockRepository.EXPECT().GetAllHosts().Return([]host.Host{host1, host2}, nil)
	mockHostsRepo.EXPECT().StoreHosts([]host.Host{host1, host2}).Return(nil)

	hostsStorage := usecase.NewHostsStorage(mockRepository, mockHostsRepo)
	err := hostsStorage.StoreHosts()

	if err != nil {
		t.Fail()
	}
}

func TestStoreHostsGetAllHostsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostsRepo := mocks.NewMockHostsRepository(ctrl)
	mockRepository := hostPortsMock.NewMockHostRepository(ctrl)
	mockRepository.EXPECT().GetAllHosts().Return([]host.Host{}, fmt.Errorf("Error test"))

	hostsStorage := usecase.NewHostsStorage(mockRepository, mockHostsRepo)
	err := hostsStorage.StoreHosts()

	if err == nil {
		t.Fail()
	}
}

func TestStoreHostsGetsErrorWhenStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostsRepo := mocks.NewMockHostsRepository(ctrl)
	mockRepository := hostPortsMock.NewMockHostRepository(ctrl)
	mockRepository.EXPECT().GetAllHosts().Return([]host.Host{host1, host2}, nil)
	mockHostsRepo.EXPECT().StoreHosts([]host.Host{host1, host2}).Return(fmt.Errorf("Error test"))

	hostsStorage := usecase.NewHostsStorage(mockRepository, mockHostsRepo)
	err := hostsStorage.StoreHosts()

	if err == nil {
		t.Fail()
	}
}
