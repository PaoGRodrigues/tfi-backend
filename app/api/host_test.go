package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateHostUseCaseAndGetAllHosts(t *testing.T) {

	var (
		name = "Test"
		ip   = "13.13.13.13"
	)
	expectedHosts := []domains.Host{
		domains.Host{
			Name: name,
			IP:   ip,
		},
	}

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostSearcherUseCase := mocks.NewMockHostUseCase(ctrl)
	mockHostSearcherUseCase.EXPECT().GetAllHosts().Return(expectedHosts, nil)

	api := &api.Api{
		HostUseCase: mockHostSearcherUseCase,
		Engine:      gin.Default(),
	}

	api.MapGetHostsURL()

	response := httptest.NewRecorder()
	requestUrl := "/hosts"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))
	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCreateAHostUsecaseAndGetHostsReturnsAnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostSearcherUseCase := mocks.NewMockHostUseCase(ctrl)
	mockHostSearcherUseCase.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing error case"))

	api := &api.Api{
		HostUseCase: mockHostSearcherUseCase,
		Engine:      gin.Default(),
	}

	api.MapGetHostsURL()

	response := httptest.NewRecorder()

	requestUrl := "/hosts"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)

}

func TestCreateHostFilterCaseAndGetAllLocalHosts(t *testing.T) {

	var (
		name        = "Test"
		ip          = "13.13.13.13"
		privateHost = true
	)

	localhosts := []domains.Host{
		domains.Host{
			Name:        name,
			IP:          ip,
			PrivateHost: privateHost,
		},
	}

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostFilter := mocks.NewMockHostsFilter(ctrl)
	mockHostFilter.EXPECT().GetLocalHosts().Return(localhosts, nil)

	api := &api.Api{
		HostsFilter: mockHostFilter,
		Engine:      gin.Default(),
	}

	api.MapGetLocalHostsURL()

	response := httptest.NewRecorder()

	requestUrl := "/localhosts"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCreateHostFilterCaseAndReturnsAnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostFilter := mocks.NewMockHostsFilter(ctrl)
	mockHostFilter.EXPECT().GetLocalHosts().Return(nil, fmt.Errorf("Testing error case"))

	api := &api.Api{
		HostsFilter: mockHostFilter,
		Engine:      gin.Default(),
	}

	api.MapGetLocalHostsURL()

	response := httptest.NewRecorder()

	requestUrl := "/localhosts"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
