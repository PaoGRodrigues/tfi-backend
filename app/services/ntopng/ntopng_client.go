package ntopng

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type NtopNG struct {
	UrlClient   string
	InterfaceId int
	Usr         string
	Pass        string
}

type Interface struct {
	IfName string `json:"ifname"`
	IfId   int    `json:"ifid"`
	Name   string `json:"name"`
}

type HTTPIfResponse struct {
	Rc      int         `json:"rc"`
	RcStrHr string      `json:"rc_str_hr"`
	RcStr   string      `json:"rc_str"`
	Rsp     []Interface `json:"rsp"`
}

func NewTool(urlClient string, usr string, pass string) *NtopNG {
	return &NtopNG{
		UrlClient: urlClient,
		Usr:       usr,
		Pass:      pass,
	}
}

func (t *NtopNG) SetInterfaceID() error {
	response, err := t.getInterfaceID()
	if err != nil {
		return err
	}
	allInterfaces := response.Rsp
	var wlan0ID int
	for _, id := range allInterfaces {
		if id.Name == "wlan0" {
			wlan0ID = id.IfId
		}
	}

	t.InterfaceId = wlan0ID
	return nil
}

func (t *NtopNG) getInterfaceID() (HTTPIfResponse, error) {
	client := &http.Client{}
	endpoint := "/lua/rest/v2/get/ntopng/interfaces.lua"

	req, err := http.NewRequest("GET", t.UrlClient+endpoint, nil)
	if err != nil {
		return HTTPIfResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(t.Usr, t.Pass)

	response, err := client.Do(req)
	if err != nil {
		return HTTPIfResponse{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return HTTPIfResponse{}, err
	}

	var resp HTTPIfResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return HTTPIfResponse{}, err
	}

	return resp, nil
}
