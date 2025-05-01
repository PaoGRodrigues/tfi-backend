package api

import (
	"fmt"
	"net/http"

	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	"github.com/gin-gonic/gin"
)

type HostsResponse struct {
	Name        string `json:"Nombre,omitempty"`
	PrivateHost bool   `json:"Privado"`
	IP          string `json:"IP"`
	Mac         string `json:"MAC"`
	ASname      string `json:"ASname,omitempty"`
}

func (api *Api) GetLocalHosts(c *gin.Context) {
	hosts, err := api.GetLocalhostsUseCase.GetLocalHosts()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": parseHostResponse(hosts)})
}

func parseHostResponse(hosts []host.Host) []HostsResponse {
	response := []HostsResponse{}

	for _, host := range hosts {
		h := HostsResponse{
			PrivateHost: host.PrivateHost,
			IP:          host.IP,
			Mac:         host.Mac,
		}
		response = append(response, h)
	}
	return response
}
