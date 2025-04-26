package host_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/usecase/host"
	hostPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/host"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var server = "13.13.13.13"

func TestBlockSourceIPReturnCorrect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlockerService := hostPortsMock.NewMockHostBlocker(ctrl)
	mockBlockerService.EXPECT().Block(server).Return(&server, nil)

	blocker := host.NewBlocker(mockBlockerService)
	get, err := blocker.Block(server)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, server, *get)
}

func TestBlockSourceIPGetHostReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlockerService := hostPortsMock.NewMockHostBlocker(ctrl)
	mockBlockerService.EXPECT().Block("sarasa").Return(nil, fmt.Errorf("Error Test"))

	blocker := host.NewBlocker(mockBlockerService)
	_, err := blocker.Block("sarasa")

	if err == nil {
		t.Fail()
	}
}

func TestBlockSourceIPBlockHostReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlockerService := hostPortsMock.NewMockHostBlocker(ctrl)
	mockBlockerService.EXPECT().Block(server).Return(nil, fmt.Errorf("Error Test"))

	blocker := host.NewBlocker(mockBlockerService)
	_, err := blocker.Block(server)

	if err == nil {
		t.Fail()
	}
}
