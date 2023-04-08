package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

// Alerts
type Alert struct {
	Row_id string `json:"-"`
	Name   struct {
		Name string `json:"fullname"`
	} `json:"Msg"`
	Family string
	Time   struct {
		Label string
	} `json:"tstamp"`
	Severity struct {
		Label string
	} `json:"label"`
	AlertFlow     AlertFlow `json:"flow"`
	AlertProtocol AlertProtocol
}

type AlertFlow struct {
	CliPort int         `json:"cli_port"`
	SrvPort int         `json:"srv_port"`
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
	} `json:"l7_proto"`
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

func (t *NtopNG) GetAllAlerts(epoch_begin, epoch_end int) ([]domains.Alert, error) {
	alertsListResponse, err := t.getAlertsList(epoch_begin, epoch_end)
	if err != nil {
		return nil, err
	}

	alerts := []domains.Alert{}
	if alertsListResponse.Rsp.Alerts != nil {

		return parseAlertsFromTool(alertsListResponse.Rsp.Alerts), nil
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

func parseAlertsFromTool(rawAlerts []Alert) []domains.Alert {

	formattedAlerts := []domains.Alert{}
	for _, alert := range rawAlerts {
		newAlert := domains.Alert{
			Name:     alert.Name.Name,
			Family:   alert.Family,
			Time:     alert.Time,
			Severity: alert.Severity.Label,
			AlertFlow: domains.AlertFlow{
				Client: flow.Client{
					Name: alert.AlertFlow.Client.Value,
					IP:   alert.AlertFlow.Client.Value,
					Port: alert.AlertFlow.CliPort,
				},
				Server: flow.Server{
					Name: alert.AlertFlow.Server.Name,
					IP:   alert.AlertFlow.Server.Value,
					Port: alert.AlertFlow.SrvPort,
				},
			},
			AlertProtocol: flow.Protocol{
				L4:    alert.AlertProtocol.Protocol.L4,
				L7:    alert.AlertProtocol.Protocol.L7,
				Label: alert.AlertProtocol.Protocol.Label,
			},
		}

		formattedAlerts = append(formattedAlerts, newAlert)
	}

	return formattedAlerts
}
