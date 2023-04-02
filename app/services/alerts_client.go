package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
)

type HttpAlertResponse struct {
	Rc              int     `json:"rc"`
	RcStrHr         string  `json:"rc_str_hr"`
	RcStr           string  `json:"rc_str"`
	Rsp             Records `json:"rsp"`
	RecordsTotal    int
	RecordsFiltered int
}

type Records struct {
	Alerts []domains.Alert `json:"records"`
}

func (t *NtopNG) GetAllAlerts(epoch_begin, epoch_end int, host string) ([]domains.Alert, error) {
	alertsListResponse, err := t.getAlertsList(epoch_begin, epoch_end, host)
	if err != nil {
		return nil, err
	}
	return alertsListResponse.Rsp.Alerts, nil
}

func (t *NtopNG) getAlertsList(epoch_begin, epoch_end int, host string) (HttpAlertResponse, error) {
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

	fmt.Print(resp)
	return resp, nil
}
