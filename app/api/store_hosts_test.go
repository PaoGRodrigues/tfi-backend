package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	hostUsecase "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"
	hostPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/host"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestStoreHostSuccessfullyReturn200(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocksHostReader := hostPortsMock.NewMockHostRepository(ctrl)
	mockHostDBRepository := hostPortsMock.NewMockHostDBRepository(ctrl)
	storeHostUsecase := hostUsecase.NewHostsStorage(mocksHostReader, mockHostDBRepository)
	mocksHostReader.EXPECT().GetAllHosts().Return([]host.Host{host1, host2}, nil)
	mockHostDBRepository.EXPECT().StoreHosts([]host.Host{host1, host2}).Return(nil)

	api := &api.Api{
		HostsStorage: storeHostUsecase,
		Engine:       gin.Default(),
	}

	api.MapStoreHostsURL()

	response := httptest.NewRecorder()

	requestUrl := "/hosts"
	httpRequest, _ := http.NewRequest("POST", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestStoreHostsFailedAndReturn500(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocksHostReader := hostPortsMock.NewMockHostRepository(ctrl)
	mockHostDBRepository := hostPortsMock.NewMockHostDBRepository(ctrl)
	mocksHostReader.EXPECT().GetAllHosts().Return([]host.Host{host1, host2}, nil)
	storeHostUsecase := hostUsecase.NewHostsStorage(mocksHostReader, mockHostDBRepository)
	mockHostDBRepository.EXPECT().StoreHosts([]host.Host{host1, host2}).Return(fmt.Errorf("Testing error case"))

	api := &api.Api{
		HostsStorage: storeHostUsecase,
		Engine:       gin.Default(),
	}

	api.MapStoreHostsURL()

	response := httptest.NewRecorder()

	requestUrl := "/hosts"
	httpRequest, _ := http.NewRequest("POST", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
