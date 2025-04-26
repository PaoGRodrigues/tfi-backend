package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	host "github.com/PaoGRodrigues/tfi-backend/app/domain/host"
)

type HttpHostResponse struct {
	Rc    int
	RcStr string
	Rsp   []host.Host
}

func (t *NtopNG) GetAllHosts() ([]host.Host, error) {

	HostsListResponse, err := t.getHostsList()
	if err != nil {
		return nil, err
	}

	return HostsListResponse.Rsp, nil
}

func (t *NtopNG) getHostsList() (HttpHostResponse, error) {
	client := &http.Client{}
	endpoint := "/lua/rest/v2/get/host/custom_data.lua"

	req, err := http.NewRequest("GET", t.UrlClient+endpoint, nil)
	if err != nil {
		return HttpHostResponse{}, err
	}
	req.SetBasicAuth(t.Usr, t.Pass)
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()

	query.Add("ifid", strconv.Itoa(t.InterfaceId))
	query.Add("field_alias", "asname,privatehost,ip,os_detail,mac,city,country")
	req.URL.RawQuery = query.Encode()

	response, err := client.Do(req)
	if err != nil {
		return HttpHostResponse{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return HttpHostResponse{}, err
	}

	var resp HttpHostResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return HttpHostResponse{}, err
	}

	return resp, nil
}
