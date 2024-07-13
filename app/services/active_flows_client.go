package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type HttpResponse struct {
	Rc    int
	RcStr string
	Rsp   RspData
}

type RspData struct {
	Data        []domains.ActiveFlow
	CurrentPage int
	PerPage     int
}

func (t *NtopNG) GetAllActiveTraffic() ([]domains.ActiveFlow, error) {
	activeFlows, err := t.getActiveFlows()
	if err != nil {
		return nil, err
	}
	return activeFlows, nil
}

func (t *NtopNG) getActiveFlows() ([]domains.ActiveFlow, error) {
	activeFlows := []domains.ActiveFlow{}
	resp, err := t.getActiveFlowsSinglePage(0)
	if err != nil {
		return nil, err
	}
	for len(resp.Rsp.Data) > resp.Rsp.PerPage {
		activeFlows = append(activeFlows, resp.Rsp.Data...)
		resp, err = t.getActiveFlowsSinglePage(resp.Rsp.CurrentPage + 1)
		if err != nil {
			return nil, err
		}
	}
	activeFlows = append(activeFlows, resp.Rsp.Data...)
	return activeFlows, nil
}

func (t *NtopNG) getActiveFlowsSinglePage(page int) (HttpResponse, error) {
	client := &http.Client{}
	endpoint := "/lua/rest/v2/get/flow/active.lua"

	req, err := http.NewRequest("GET", t.UrlClient+endpoint, nil)
	if err != nil {
		return HttpResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(t.Usr, t.Pass)

	query := req.URL.Query()
	query.Add("ifid", strconv.Itoa(t.InterfaceId))
	query.Add("currentPage", strconv.Itoa(page))

	req.URL.RawQuery = query.Encode()

	response, err := client.Do(req)
	if err != nil {
		return HttpResponse{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return HttpResponse{}, err
	}

	var resp HttpResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return HttpResponse{}, err
	}
	return resp, nil
}
