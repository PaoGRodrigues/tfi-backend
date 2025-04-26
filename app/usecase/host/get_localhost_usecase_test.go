package host_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	usecase "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
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

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return(expected)

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	got, err := filter.GetLocalHosts()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, []host.Host{local}, got)
}

func TestGetLocalHostCallingGetHostFromRepoInSearcherReturnLocalHosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return([]host.Host{})
	mockSearcher.EXPECT().GetAllHosts().Return(expected, nil)

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	got, err := filter.GetLocalHosts()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, []host.Host{local}, got)
}

func TestGetLocalHostAndGetAllHostsInSearcherReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return([]host.Host{})
	mockSearcher.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing Error"))

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
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

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	got, err := filter.GetRemoteHosts()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, []host.Host{remote}, got)
}

func TestGetRemoteHostCallingGetHostFromRepoInSearcherReturnRemoteHosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return([]host.Host{})
	mockSearcher.EXPECT().GetAllHosts().Return(expected, nil)

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	got, err := filter.GetRemoteHosts()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, []host.Host{remote}, got)
}

func TestGetRemoteHostAndGetAllHostsInSearcherReturnAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return([]host.Host{})
	mockSearcher.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing Error"))

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	_, err := filter.GetRemoteHosts()

	if err == nil {
		t.Fail()
	}
}

func TestGetHostByIPReturnCorrectHost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return(expected)

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	got, err := filter.GetHost(local.IP)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, local, got)
}

func TestGetHostByIPGetAllHostsReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return(nil)
	mockSearcher.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Error Test"))

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	_, err := filter.GetHost(local.IP)

	if err == nil {
		t.Fail()
	}
}

func TestGetHostByIPWithAnUnexistingIPReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return(expected)

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	_, err := filter.GetHost("10.10.10.10")

	if err == nil {
		t.Fail()
	}
}

func TestGetHostByURLReturnCorrectHost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return(expected)

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	got, err := filter.GetHost(remote.Name)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, remote, got)
}

func TestGetHostByAnUnexistingURLReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockHostUseCase(ctrl)
	mockSearcher.EXPECT().GetHosts().Return(expected)

	filter := usecase.NewGetLocalhostsUseCase(mockSearcher)
	_, err := filter.GetHost("test.test.com")

	if err == nil {
		t.Fail()
	}
}
