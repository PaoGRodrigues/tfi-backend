package usecase_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var hosts = []domains.Host{
	{
		Name:        "Test",
		IP:          "13.13.13.13",
		PrivateHost: true,
	},
	{
		Name:        "Test2",
		IP:          "172.172.172.172",
		PrivateHost: false,
	},
}

func TestBlockSourceIPReturnCorrectHost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilter := mocks.NewMockHostsFilter(ctrl)
	mockFilter.EXPECT().GetHost(hosts[0].IP).Return(hosts[0], nil)

	mockBlockerService := mocks.NewMockHostBlockService(ctrl)
	mockBlockerService.EXPECT().BlockHost(hosts[0]).Return(nil)

	blocker := usecase.NewBlocker(mockBlockerService, mockFilter)
	get, err := blocker.Block(hosts[0].IP)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, hosts[0], get)
}

func TestBlockSourceIPGetHostReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilter := mocks.NewMockHostsFilter(ctrl)
	mockFilter.EXPECT().GetHost(hosts[0].IP).Return(domains.Host{}, fmt.Errorf("Error Test"))

	mockBlockerService := mocks.NewMockHostBlockService(ctrl)

	blocker := usecase.NewBlocker(mockBlockerService, mockFilter)
	_, err := blocker.Block(hosts[0].IP)

	if err == nil {
		t.Fail()
	}
}

func TestBlockSourceIPBlockHostReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilter := mocks.NewMockHostsFilter(ctrl)
	mockFilter.EXPECT().GetHost(hosts[0].IP).Return(hosts[0], nil)

	mockBlockerService := mocks.NewMockHostBlockService(ctrl)
	mockBlockerService.EXPECT().BlockHost(hosts[0]).Return(fmt.Errorf("Error Test"))

	blocker := usecase.NewBlocker(mockBlockerService, mockFilter)
	_, err := blocker.Block(hosts[0].IP)

	if err == nil {
		t.Fail()
	}
}
