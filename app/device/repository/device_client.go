package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
	tool "github.com/PaoGRodrigues/tfi-backend/app/services/tool"
)

type DeviceClient struct {
	tool     *tool.Tool
	endpoint string
}
type HttpResponse struct {
	Rc    int
	RcStr string
	Rsp   []domains.Device
}

func NewDeviceClient(tool *tool.Tool, endpoint string) *DeviceClient {

	return &DeviceClient{
		tool:     tool,
		endpoint: endpoint,
	}
}

func (d *DeviceClient) GetAll() ([]domains.Device, error) {

	devicesListResponse, err := d.getDevicesList()
	if err != nil {
		return nil, err
	}

	return devicesListResponse.Rsp, nil
}

func (d *DeviceClient) getDevicesList() (HttpResponse, error) {
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

	fmt.Printf(string(body))
	var resp HttpResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return HttpResponse{}, err
	}

	return resp, nil
}
