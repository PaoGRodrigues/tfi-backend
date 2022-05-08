package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/host/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/host"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateHostUseCaseAndGetAllHosts(t *testing.T) {

	var (
		name = "Test"
		ip   = "13.13.13.13"
	)

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostSearcherUseCase := mocks.NewMockHostUseCase(ctrl)

	api := &api.Api{
		HostUseCase: mockHostSearcherUseCase,
		Engine:      gin.Default(),
	}

	r := gin.Default()

	r.GET("/hosts", api.GetHosts,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/hosts"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	createdHosts := []domains.Host{
		domains.Host{
			Name: name,
			IP:   ip,
		},
	}

	t.Run("Ok", func(t *testing.T) {

		mockHostSearcherUseCase.EXPECT().GetAllHosts().Return(createdHosts, nil)

		res := executeWithContext()
		assert.Equal(t, http.StatusOK, res.Code)
	})
}

func TestCreateAHostUsecaseAndGetHostsReturnsAnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostSearcherUseCase := mocks.NewMockHostUseCase(ctrl)

	api := &api.Api{
		HostUseCase: mockHostSearcherUseCase,
		Engine:      gin.Default(),
	}

	r := gin.Default()

	r.GET("/hosts", api.GetHosts,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/hosts"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	t.Run("Ok", func(t *testing.T) {

		mockHostSearcherUseCase.EXPECT().GetAllHosts().Return(nil, fmt.Errorf("Testing error case"))

		res := executeWithContext()
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestCreateHostFilterCaseAndGetAllLocalHosts(t *testing.T) {

	var (
		name        = "Test"
		ip          = "13.13.13.13"
		privateHost = true
	)

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostFilter := mocks.NewMockHostsFilter(ctrl)

	api := &api.Api{
		HostsFilter: mockHostFilter,
		Engine:      gin.Default(),
	}

	r := gin.Default()

	r.GET("/localhosts", api.GetLocalHosts,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/localhosts"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	localhosts := []domains.Host{
		domains.Host{
			Name:        name,
			IP:          ip,
			PrivateHost: privateHost,
		},
	}

	t.Run("Ok", func(t *testing.T) {

		mockHostFilter.EXPECT().GetLocalHosts().Return(localhosts, nil)

		res := executeWithContext()
		assert.Equal(t, http.StatusOK, res.Code)
	})
}

func TestCreateHostFilterCaseAndReturnsAnError(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHostFilter := mocks.NewMockHostsFilter(ctrl)

	api := &api.Api{
		HostsFilter: mockHostFilter,
		Engine:      gin.Default(),
	}

	r := gin.Default()

	r.GET("/localhosts", api.GetLocalHosts,
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	executeWithContext := func() *httptest.ResponseRecorder {
		response := httptest.NewRecorder()

		requestUrl := "/localhosts"
		httpRequest, _ := http.NewRequest("GET", requestUrl, strings.NewReader(string("")))

		r.ServeHTTP(response, httpRequest)
		return response
	}

	t.Run("Ok", func(t *testing.T) {

		mockHostFilter.EXPECT().GetLocalHosts().Return(nil, fmt.Errorf("Testing error case"))

		res := executeWithContext()
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
