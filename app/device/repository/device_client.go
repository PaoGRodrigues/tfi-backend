package repository

import (
	"encoding/json"
	"fmt"
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
	Rc    int
	RcStr string
	Rsp   []domains.Device
}

func NewDeviceClient(urlClient string, interfaceId int, usr string, pass string) *DeviceClient {

	return &DeviceClient{
		urlClient:   urlClient,
		interfaceId: interfaceId,
		usr:         usr,
		pass:        pass,
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
	endpoint := "/lua/rest/v2/get/host/custom_data.lua"

	req, err := http.NewRequest("GET", d.urlClient+endpoint, nil)
	if err != nil {
		return HttpResponse{}, err
	}
	req.SetBasicAuth(d.usr, d.pass)
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("ifid", string("2"))
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
