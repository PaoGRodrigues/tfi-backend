package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PaoGRodrigues/tfi-backend/app/device/domains"
)

const URL = "http://localhost:3000"

type DeviceClient struct {
	endpoint    string
	interfaceId int
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

func NewDeviceClient(endpoint string, interfaceId int) *DeviceClient {

	return &DeviceClient{
		endpoint:    endpoint,
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
	uri := "/lua/rest/v2/get/host/active.lua"

	u, err := url.Parse(d.endpoint + uri)
	if err != nil {
		return HttpResponse{}, err
	}
	query := u.Query()
	query.Add("ifid", string(d.interfaceId))
	u.RawQuery = query.Encode()

	response, err := http.Get(u.String())
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
