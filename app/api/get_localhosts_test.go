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

func TestCreateHostFilterCaseAndGetAllLocalHosts(t *testing.T) {

	var (
		name        = "Test"
		ip          = "13.13.13.13"
		privateHost = true
	)

	localhosts := []host.Host{
		host.Host{
			Name:        name,
			IP:          ip,
			PrivateHost: privateHost,
		},
	}

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := hostPortsMock.NewMockHostRepository(ctrl)
	getLocalhostUseCase := hostUsecase.NewGetLocalhostsUseCase(mockRepository)
	mockRepository.EXPECT().GetAllHosts().Return(localhosts, nil)

	api := &api.Api{
		GetLocalhostsUseCase: getLocalhostUseCase,
		Engine:               gin.Default(),
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

	mockRepository := hostPortsMock.NewMockHostRepository(ctrl)
	getLocalhostUseCase := hostUsecase.NewGetLocalhostsUseCase(mockRepository)
	mockRepository.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing error case"))

	api := &api.Api{
		GetLocalhostsUseCase: getLocalhostUseCase,
		Engine:               gin.Default(),
	}

	api.MapGetLocalHostsURL()

	response := httptest.NewRecorder()

	requestUrl := "/localhosts"
	httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
