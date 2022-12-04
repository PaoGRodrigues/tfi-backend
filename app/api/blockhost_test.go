package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var Host = domains.Host{
	Name: "test.google.com",
	IP:   "192.192.192.10",
}

type blockHostRequest struct {
	Host string // Host can be IP or URL
}

func TestBlockHostByIPReturn200(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlocker := mocks.NewMockHostBlocker(ctrl)
	mockBlocker.EXPECT().Block(Host.IP).Return(Host, nil)

	api := &api.Api{
		HostBlocker: mockBlocker,
		Engine:      gin.Default(),
	}

	api.MapBlockHost()

	response := httptest.NewRecorder()

	req := blockHostRequest{
		Host: "192.192.192.10",
	}

	body, _ := json.Marshal(req)

	requestUrl := "/blockedhosts"
	httpRequest, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))

	api.Engine.ServeHTTP(response, httpRequest)

	assert.Equal(t, http.StatusOK, response.Code)
}
