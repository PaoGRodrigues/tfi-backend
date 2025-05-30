package host_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	usecase "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"
	hostPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/host"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var local = host.Host{
	Name:        "Test",
	IP:          "13.13.13.13",
	PrivateHost: true,
}
var remote = host.Host{
	Name:        "Test2.google.com",
	IP:          "172.172.172.172",
	PrivateHost: false,
}

var expected = []host.Host{
	local,
	remote,
}

func TestGetLocalHostWithHostsReturnedFromSearcherReturnLocalHosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := hostPortsMock.NewMockHostRepository(ctrl)
	mockRepository.EXPECT().GetAllHosts().Return(expected, nil)

	filter := usecase.NewGetLocalhostsUseCase(mockRepository)
	got, err := filter.GetLocalHosts()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, []host.Host{local}, got)
}

func TestGetLocalHostAndGetAllHostsInSearcherReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := hostPortsMock.NewMockHostRepository(ctrl)
	mockRepository.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing Error"))

	filter := usecase.NewGetLocalhostsUseCase(mockRepository)
	_, err := filter.GetLocalHosts()

	if err == nil {
		t.Fail()
	}
}
