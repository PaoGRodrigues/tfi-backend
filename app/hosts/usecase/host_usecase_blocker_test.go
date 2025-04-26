package usecase_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/hosts/usecase"
	hosts_mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var server = "13.13.13.13"

func TestBlockSourceIPReturnCorrect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlockerService := hosts_mocks.NewMockHostBlockerService(ctrl)
	mockBlockerService.EXPECT().BlockHost(server).Return(nil)

	blocker := usecase.NewBlocker(mockBlockerService)
	get, err := blocker.Block(server)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, server, *get)
}

func TestBlockSourceIPGetHostReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlockerService := hosts_mocks.NewMockHostBlockerService(ctrl)
	mockBlockerService.EXPECT().BlockHost("sarasa").Return(fmt.Errorf("Error Test"))

	blocker := usecase.NewBlocker(mockBlockerService)
	_, err := blocker.Block("sarasa")

	if err == nil {
		t.Fail()
	}
}

func TestBlockSourceIPBlockHostReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlockerService := hosts_mocks.NewMockHostBlockerService(ctrl)
	mockBlockerService.EXPECT().BlockHost(server).Return(fmt.Errorf("Error Test"))

	blocker := usecase.NewBlocker(mockBlockerService)
	_, err := blocker.Block(server)

	if err == nil {
		t.Fail()
	}
}
