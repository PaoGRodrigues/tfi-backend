package ntopng

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	flow "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
)

// Alerts
type Alert struct {
	Row_id string `json:"-"`
	Name   struct {
		Name string `json:"fullname"`
	} `json:"Msg"`
	Family   string
	Category struct {
		Label string
	} `json:"alert_category"`
	Time struct {
		Label string
	} `json:"tstamp"`
	Severity struct {
		Value int
	} `json:"severity"`
	AlertFlow     AlertFlow     `json:"flow"`
	AlertProtocol AlertProtocol `json:"l7_proto"`
}

type AlertFlow struct {
	CliPort string      `json:"cli_port"`
	SrvPort string      `json:"srv_port"`
	Client  AlertClient `json:"cli_ip"`
	Server  AlertServer `json:"srv_ip"`
}

type AlertClient struct {
	Value  string `json:"value"`
	Contry string `json:"country"`
}

type AlertServer struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Country string `json:"country"`
}

type AlertProtocol struct {
	Protocol struct {
		L4    string `json:"l4_label"`
		Label string `json:"label"`
		L7    string `json:"l7_label"`
	}
}

type HttpAlertResponse struct {
	Rc              int     `json:"rc"`
	RcStrHr         string  `json:"rc_str_hr"`
	RcStr           string  `json:"rc_str"`
	Rsp             Records `json:"rsp"`
	RecordsTotal    int
	RecordsFiltered int
}

type Records struct {
	Alerts []Alert `json:"records"`
}

var severityScore = map[int]string{
	1: "Depuración",
	2: "Informativo",
	3: "Notificación",
	4: "Advertencia",
	5: "Error",
	6: "Crítico",
	7: "Alerta",
	8: "Emergencia",
}

func (t *NtopNG) GetAllAlerts(epoch_begin, epoch_end int) ([]alert.Alert, error) {
	alertsListResponse, err := t.getAlertsList(epoch_begin, epoch_end)
	if err != nil {
		return nil, err
	}

	alerts := []alert.Alert{}
	if alertsListResponse.Rsp.Alerts != nil {
		parsedAlerts, err := parseAlertsFromTool(alertsListResponse.Rsp.Alerts)
		if err != nil {
			return nil, err
		}
		return parsedAlerts, nil
	}

	return alerts, nil
}

func (t *NtopNG) getAlertsList(epoch_begin, epoch_end int) (HttpAlertResponse, error) {
	client := &http.Client{}
	endpoint := "/lua/rest/v2/get/flow/alert/list.lua"

	req, err := http.NewRequest("GET", t.UrlClient+endpoint, nil)
	if err != nil {
		return HttpAlertResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(t.Usr, t.Pass)

	query := req.URL.Query()
	query.Add("ifid", strconv.Itoa(t.InterfaceId))
	query.Add("epoch_begin", strconv.Itoa(epoch_begin))
	query.Add("epoch_end", strconv.Itoa(epoch_end))

	req.URL.RawQuery = query.Encode()

	response, err := client.Do(req)
	if err != nil {
		return HttpAlertResponse{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return HttpAlertResponse{}, err
	}

	var resp HttpAlertResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return HttpAlertResponse{}, err
	}

	return resp, nil
}

func parseAlertsFromTool(rawAlerts []Alert) ([]alert.Alert, error) {

	formattedAlerts := []alert.Alert{}
	for _, anAlert := range rawAlerts {
		cliPort, err := strconv.Atoi(anAlert.AlertFlow.CliPort)
		if err != nil {
			return formattedAlerts, err
		}
		srvPort, err := strconv.Atoi(anAlert.AlertFlow.SrvPort)
		if err != nil {
			return formattedAlerts, err
		}

		newAlert := alert.Alert{
			Name:     anAlert.Name.Name,
			Family:   anAlert.Family,
			Category: anAlert.Category.Label,
			Time:     anAlert.Time.Label,
			Severity: severityScore[anAlert.Severity.Value],
			AlertFlow: alert.AlertFlow{
				Client: flow.Client{
					Name: anAlert.AlertFlow.Client.Value,
					IP:   anAlert.AlertFlow.Client.Value,
					Port: cliPort,
				},
				Server: flow.Server{
					Name: anAlert.AlertFlow.Server.Name,
					IP:   anAlert.AlertFlow.Server.Value,
					Port: srvPort,
				},
			},
			AlertProtocol: flow.Protocol{
				L4:    anAlert.AlertProtocol.Protocol.L4,
				L7:    anAlert.AlertProtocol.Protocol.L7,
				Label: anAlert.AlertProtocol.Protocol.Label,
			},
		}

		formattedAlerts = append(formattedAlerts, newAlert)
	}

	return formattedAlerts, nil
}
