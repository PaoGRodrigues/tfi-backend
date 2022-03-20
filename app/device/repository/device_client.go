package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
)

const URL = "http://192.168.0.16:3000"

type DeviceClient struct {
	urlClient   string
	interfaceId int
	usr         string
	pass        string
}
type HttpResponse struct {
	Rc    string
	RcStr string
	Rsp   response
}
type response struct {
	CurrentPage int
	Data        []data
}

type data struct {
	IsLocalhost bool `json:is_localhost"`
	Country     string
	Name        string
	IP          string
	OsDetail    string `json:"os_detail"`
	Mac         string
}

func NewDeviceClient(urlClient string, interfaceId int) *DeviceClient {

	return &DeviceClient{
		urlClient:   urlClient,
		interfaceId: interfaceId,
	}
}

func (d *DeviceClient) GetAll() ([]domains.Device, error) {

	devicesListResponse, err := d.getDevicesList()
	if err != nil {
		return nil, err
	}

	devices := []domains.Device{}
	for _, dev := range devicesListResponse.Rsp.Data {

	}

	return devices, nil
}

func (d *DeviceClient) getDevicesList() (HttpResponse, error) {
	client := &http.Client{}
	endpoint := "/lua/rest/v2/get/host/custom_data.lua"

	req, err := http.NewRequest("GET", d.urlClient+endpoint, nil)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(d.usr, d.pass)

	query := req.URL.Query()
	query.Add("ifid", string(d.interfaceId))

	req.URL.RawQuery = query.Encode()

	response, _ := client.Do(req)
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

// -d '{"ifid": 2, "field_alias": "is_localhost,name,privatehost,ip,os_detail,mac,city,country"}'
