package usecase_test

import (
	"fmt"
	"testing"

	host_domains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/usecase"
	traffic_domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	hosts_mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	traffic_mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var servers = []traffic_domains.Server{
	{
		Name: "Test",
		IP:   "13.13.13.13",
	},
	{
		Name: "Test2",
		IP:   "172.172.172.172",
	},
}

func TestBlockSourceIPReturnCorrectHost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilter := traffic_mocks.NewMockTrafficRepository(ctrl)
	mockFilter.EXPECT().GetServerByAttr(servers[0].IP).Return(servers[0], nil)

	host := host_domains.Host{IP: servers[0].IP, Name: servers[0].Name}

	mockBlockerService := hosts_mocks.NewMockHostBlockerService(ctrl)
	mockBlockerService.EXPECT().BlockHost(host).Return(nil)

	blocker := usecase.NewBlocker(mockBlockerService, mockFilter)
	get, err := blocker.Block(servers[0].IP)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, host, get)
}

func TestBlockSourceIPGetHostReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilter := traffic_mocks.NewMockTrafficRepository(ctrl)
	mockFilter.EXPECT().GetServerByAttr(servers[0].IP).Return(traffic_domains.Server{}, fmt.Errorf("Error Test"))

	mockBlockerService := hosts_mocks.NewMockHostBlockerService(ctrl)

	blocker := usecase.NewBlocker(mockBlockerService, mockFilter)
	_, err := blocker.Block(servers[0].IP)

	if err == nil {
		t.Fail()
	}
}

func TestBlockSourceIPBlockHostReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilter := traffic_mocks.NewMockTrafficRepository(ctrl)
	mockFilter.EXPECT().GetServerByAttr(servers[0].IP).Return(servers[0], nil)

	host := host_domains.Host{IP: servers[0].IP, Name: servers[0].Name}

	mockBlockerService := hosts_mocks.NewMockHostBlockerService(ctrl)
	mockBlockerService.EXPECT().BlockHost(host).Return(fmt.Errorf("Error Test"))

	blocker := usecase.NewBlocker(mockBlockerService, mockFilter)
	_, err := blocker.Block(servers[0].IP)

	if err == nil {
		t.Fail()
	}
}
