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

var host1 = domains.Host{
	Name:        "test",
	PrivateHost: false,
	IP:          "123.123.123.123",
	City:        "",
	Country:     "US",
}

var host2 = domains.Host{
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

	mockHostStorage := mocks.NewMockHostsStorage(ctrl)
	mockHostStorage.EXPECT().StoreHosts().Return(nil)

	api := &api.Api{
		HostsStorage: mockHostStorage,
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

	mockHostStorage := mocks.NewMockHostsStorage(ctrl)
	mockHostStorage.EXPECT().StoreHosts().Return(fmt.Errorf("Testing error case"))

	api := &api.Api{
		HostsStorage: mockHostStorage,
		Engine:       gin.Default(),
	}

	api.MapStoreHostsURL()

	response := httptest.NewRecorder()

	requestUrl := "/hosts"
	httpRequest, _ := http.NewRequest("POST", requestUrl, strings.NewReader(string("")))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
