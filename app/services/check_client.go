package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpCheckResponse struct {
	Rc      int
	RcStrHr string
	Rsp     Response
}
type Response struct {
	Success bool
}

func (t *NtopNG) EnableChecks() {
	for _, check := range Checks {
		rsp, err := t.enableCheck(check)
		if !rsp.Rsp.Success {
			log.Fatalf("Error enabling check %s", err)
		}
		continue
	}
}

func (t *NtopNG) enableCheck(currenrCheck Check) (HttpCheckResponse, error) {
	client := &http.Client{}
	endpoint := "/lua/rest/v2/enable/check.lua"

	marshalled, err := json.Marshal(currenrCheck)
	if err != nil {
		return HttpCheckResponse{}, err
	}
	req, err := http.NewRequest("POST", t.UrlClient+endpoint, bytes.NewReader(marshalled))
	if err != nil {
		return HttpCheckResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(t.Usr, t.Pass)

	response, err := client.Do(req)
	if err != nil {
		return HttpCheckResponse{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return HttpCheckResponse{}, err
	}

	var resp HttpCheckResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return HttpCheckResponse{}, err
	}
	return resp, nil

}
