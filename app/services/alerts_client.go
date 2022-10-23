package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
)

type HttpAlertResponse struct {
	Rc      int
	RcStrHr string
	RcStr   string
	Rsp     Records
}

type Records struct {
	alerts []domains.Alert
}

func (t *NtopNG) GetAllAlerts(epoch_begin, epoch_end int) ([]domains.Alert, error) {
	alertsListResponse, err := t.getAlertsList(epoch_begin, epoch_end)
	if err != nil {
		return nil, err
	}
	return alertsListResponse.Rsp.alerts, nil
}

func (t *NtopNG) getAlertsList(epoch_begin, epoch_end int) (HttpAlertResponse, error) {
	client := &http.Client{}
	endpoint := "/lua/rest/v2/get/all/alert/list.lua"

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
