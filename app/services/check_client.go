package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
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
	checksChannel := make(chan Check)

	// Put data into the initial channel
	go func() {
		for _, check := range Checks {
			checksChannel <- check
		}
		close(checksChannel)
	}()
	fmt.Print(len(checksChannel))

	// Declare the worker function which reads de checkChannel to make the request
	worker := func() {
		for check := range checksChannel {
			rsp, err := t.enableCheck(check)
			fmt.Printf("Enabling %s", check)
			fmt.Print(rsp.Rsp.Success)
			if err != nil {
				log.Fatalf("Error enabling check %s", check)
			}
			if !rsp.Rsp.Success {
				log.Fatalf("Error enabling check %s", check)
			}
		}
	}

	// WaitGroup is used for synchronous closing of the results channel when all work is done
	numWorkers := 5
	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			worker()
			wg.Done()
		}()
	}
}

func (t *NtopNG) enableCheck(currentCheck Check) (HttpCheckResponse, error) {
	client := &http.Client{}
	endpoint := "/lua/rest/v2/enable/check.lua"

	marshalled, err := json.Marshal(currentCheck)
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
