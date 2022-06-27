package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PaoGRodrigues/tfi-backend/app/host/domains"
	services "github.com/PaoGRodrigues/tfi-backend/app/services"
)

type HostClient struct {
	tool     *services.Tool
	endpoint string
}
type HttpResponse struct {
	Rc    int
	RcStr string
	Rsp   []domains.Host
}

func NewHostClient(tool *services.Tool, endpoint string) *HostClient {

	return &HostClient{
		tool:     tool,
		endpoint: endpoint,
	}
}

func (d *HostClient) GetAll() ([]domains.Host, error) {

	HostsListResponse, err := d.getHostsList()
	if err != nil {
		return nil, err
	}

	return HostsListResponse.Rsp, nil
}

func (d *HostClient) getHostsList() (HttpResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", d.tool.UrlClient+d.endpoint, nil)
	if err != nil {
		return HttpResponse{}, err
	}
	req.SetBasicAuth(d.tool.Usr, d.tool.Pass)
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()

	query.Add("ifid", strconv.Itoa(d.tool.InterfaceId))
	query.Add("field_alias", "name,privatehost,ip,os_detail,mac,city,country")
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
