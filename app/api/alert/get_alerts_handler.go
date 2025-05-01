package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	"github.com/gin-gonic/gin"
)

type AlertsResponse struct {
	Name        string `json:"Nombre"`
	Category    string `json:"Categoría"`
	Time        string `json:"Fecha/hora"`
	Severity    string `json:"Severidad"`
	Source      string `json:"Origen"`
	Destination string `json:"Destino"`
}

func (api *Api) GetAlerts(c *gin.Context) {
	alerts, err := api.GetAlertsUseCase.GetAllAlerts()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	alertsResponse := api.parseAlertsData(alerts)
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": alertsResponse})
}

func (api *Api) parseAlertsData(alerts []alert.Alert) []AlertsResponse {
	response := []AlertsResponse{}
	for _, alert := range alerts {
		source, destination := createFlowString(alert.AlertFlow)
		ar := AlertsResponse{
			Name:        alert.Name,
			Time:        alert.Time,
			Category:    alert.Category,
			Severity:    alert.Severity,
			Source:      source,
			Destination: destination,
		}
		response = append(response, ar)
	}
	return response
}

func createFlowString(flow alert.AlertFlow) (string, string) {
	var source strings.Builder
	var destination strings.Builder

	destName := flow.Server.Name

	source.WriteString(flow.Client.IP)
	source.WriteString(":")
	source.WriteString(strconv.Itoa(flow.Client.Port))
	if destName == "" {
		destName = flow.Server.IP
	}
	destination.WriteString(destName)
	destination.WriteString(":")
	destination.WriteString(strconv.Itoa(flow.Server.Port))

	return source.String(), destination.String()
}
